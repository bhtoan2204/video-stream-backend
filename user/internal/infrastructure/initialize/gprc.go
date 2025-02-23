package initialize

import (
	"fmt"
	"log"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/infrastructure/grpc/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

// ConsulResolver implements gRPC resolver using Consul
type ConsulResolver struct {
	cc       resolver.ClientConn
	doneChan chan struct{}
}

// NewConsulResolver creates a new resolver
func NewConsulResolver(cc resolver.ClientConn) *ConsulResolver {
	r := &ConsulResolver{
		cc:       cc,
		doneChan: make(chan struct{}),
	}
	go r.watch()
	return r
}

// ResolveNow updates addresses from Consul
func (r *ConsulResolver) ResolveNow(resolver.ResolveNowOptions) {
	r.updateAddresses()
}

// Close stops the resolver
func (r *ConsulResolver) Close() {
	close(r.doneChan)
}

// updateAddresses fetches service instances from Consul
func (r *ConsulResolver) updateAddresses() {
	services, _, err := global.ConsulClient.Health().Service("users-service", "", true, nil)
	if err != nil {
		log.Printf("Failed to fetch services from Consul: %v", err)
		return
	}

	var addresses []resolver.Address
	for _, service := range services {
		addr := fmt.Sprintf("%s:%d", service.Service.Address, service.Service.Port)
		addresses = append(addresses, resolver.Address{Addr: addr})
	}

	r.cc.UpdateState(resolver.State{Addresses: addresses})
}

// watch periodically updates addresses
func (r *ConsulResolver) watch() {
	for {
		select {
		case <-r.doneChan:
			return
		default:
			r.updateAddresses()
		}
	}
}

// ConsulBuilder is the resolver builder
type ConsulBuilder struct{}

// Build creates a new resolver
func (b *ConsulBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	return NewConsulResolver(cc), nil
}

// Scheme returns the resolver scheme
func (b *ConsulBuilder) Scheme() string {
	return "consul"
}

// Register the resolver
func init() {
	resolver.Register(&ConsulBuilder{})
}

// InitGrpcClient initializes the gRPC client using Consul
func InitGrpcClient() {
	// Connect to users-service via Consul
	conn, err := grpc.Dial(
		"consul:///user-service",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatalf("Failed to connect to User Service via Consul: %v", err)
	}

	// Store gRPC client in global variable
	global.UserGrpcClient = user.NewUserServiceClient(conn)
	log.Println("gRPC Client for users-service initialized successfully")
}
