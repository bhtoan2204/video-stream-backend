package initialize

import (
	"fmt"
	"net"
	"user/global"

	"github.com/hashicorp/consul/api"
)

func InitConsul() {
	globalConsulConfig := global.Config.ConsulConfig

	consulConfig := api.DefaultConfig()
	consulConfig.Address = globalConsulConfig.Address
	consulConfig.Scheme = globalConsulConfig.Scheme
	consulConfig.Datacenter = globalConsulConfig.DataCenter
	consulConfig.Token = globalConsulConfig.Token

	registration := &api.AgentServiceRegistration{
		ID:      "user-service",
		Name:    "user-service",
		Address: consulConfig.Address,
		Port:    global.Listener.Addr().(*net.TCPAddr).Port,
		Tags:    []string{"api", "user"},
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/api/v1/health", global.Listener.Addr().(*net.TCPAddr).IP, global.Listener.Addr().(*net.TCPAddr).Port),
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
}
