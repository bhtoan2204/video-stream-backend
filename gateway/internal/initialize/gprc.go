package initialize

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/pkg/grpc/proto/user"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type ConsulResolver struct {
	consulClient *api.Client
	serviceName  string
	cc           resolver.ClientConn
	ctx          context.Context
	cancel       context.CancelFunc
	wg           sync.WaitGroup
}

func (r *ConsulResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.cc = cc
	r.ctx, r.cancel = context.WithCancel(context.Background())
	r.wg.Add(1)
	go r.watch()
	return r, nil
}

func (r *ConsulResolver) Scheme() string {
	return "consul"
}

func (r *ConsulResolver) ResolveNow(o resolver.ResolveNowOptions) {}

func (r *ConsulResolver) Close() {
	r.cancel()
	r.wg.Wait()
}

func (r *ConsulResolver) watch() {
	defer r.wg.Done()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.ctx.Done():
			return
		case <-ticker.C:
			addresses, err := r.getAddresses()
			if err != nil {
				fmt.Printf("Failed to get addresses: %v\n", err)
				continue
			}
			r.updateState(addresses)
		}
	}
}

func (r *ConsulResolver) getAddresses() ([]resolver.Address, error) {
	services, _, err := r.consulClient.Health().Service(r.serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}

	var addresses []resolver.Address
	for _, serviceEntry := range services {
		svc := serviceEntry.Service
		host := svc.Address
		if host == "" {
			host = serviceEntry.Node.Address
		}
		if strings.Contains(host, ":") {
			var err error
			host, _, err = net.SplitHostPort(host)
			if err != nil {
				return nil, err
			}
		}
		address := fmt.Sprintf("%s:%d", host, svc.Port)
		addresses = append(addresses, resolver.Address{Addr: address})
	}
	fmt.Printf("Resolved addresses: %v\n", addresses)
	return addresses, nil
}

func (r *ConsulResolver) updateState(addresses []resolver.Address) {
	state := resolver.State{Addresses: addresses}
	r.cc.UpdateState(state)
}

func InitUserGRPCClient() {
	resolver.Register(&ConsulResolver{
		consulClient: global.ConsulClient,
		serviceName:  "user-service",
	})

	conn, err := grpc.NewClient(
		"consul:///user-service", // scheme "consul"
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)

	if err != nil {
		global.Logger.Error("Failed to connect to user grpc server", zap.Error(err))
		panic(err)
	}

	global.UserGRPCClient = user.NewUserServiceClient(conn)
	global.Logger.Info("Initialize user grpc client successfully")
}
