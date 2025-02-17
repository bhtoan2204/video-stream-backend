package initialize

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bhtoan2204/video/global"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

func InitConsul() {
	globalConsulConfig := global.Config.ConsulConfig

	consulConfig := api.DefaultConfig()
	consulConfig.Address = globalConsulConfig.Address
	consulConfig.Scheme = globalConsulConfig.Scheme
	consulConfig.Datacenter = globalConsulConfig.DataCenter
	consulConfig.Token = globalConsulConfig.Token

	serviceID := uuid.New().String()
	serviceAddress := global.Listener.Addr().(*net.TCPAddr).IP.String()
	servicePort := global.Listener.Addr().(*net.TCPAddr).Port

	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "video-service",
		Address: consulConfig.Address,
		Port:    global.Listener.Addr().(*net.TCPAddr).Port,
		Tags:    []string{"api", "video"},
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/api/v1/health", serviceAddress, servicePort),
			Interval: "10s",
			Timeout:  "5s",
		},
	}
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		// global.Logger.Error("Failed to connect to Consul:", zap.Error(err))
		panic(err)
	}
	global.ConsulClient = consulClient

	err = global.ConsulClient.Agent().ServiceRegister(registration)
	if err != nil {
		// global.Logger.Error("Failed to register service:", zap.Error(err))
		panic(err)
	}
	go handleShutdown(serviceID)
}

func handleShutdown(serviceID string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	// global.Logger.Info("Shutting down service, unregistering from Consul...", zap.String("serviceID", serviceID))

	err := global.ConsulClient.Agent().ServiceDeregister(serviceID)
	if err != nil {
		// global.Logger.Error("Failed to unregister service from Consul", zap.Error(err))
		fmt.Print("Failed to unregister service from Consul")
	} else {
		// global.Logger.Info("Service unregistered successfully from Consul")
		fmt.Print("Service unregistered successfully from Consul")
	}

	os.Exit(0)
}
