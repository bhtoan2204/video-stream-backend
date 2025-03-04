package initialize

import (
	"fmt"
	"net"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/utils"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

func InitConsul() func() {
	globalConsulConfig := global.Config.ConsulConfig

	consulConfig := api.DefaultConfig()
	consulConfig.Address = globalConsulConfig.Address
	consulConfig.Scheme = globalConsulConfig.Scheme
	consulConfig.Datacenter = globalConsulConfig.DataCenter
	consulConfig.Token = globalConsulConfig.Token

	serviceID := uuid.New().String()
	servicePort := global.Listener.Addr().(*net.TCPAddr).Port
	serviceAddress, err := utils.GetInternalIP()
	if err != nil {
		global.Logger.Error("Failed to get internal IP address:", zap.Error(err))
		panic(err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "user-service",
		Address: serviceAddress,
		Port:    servicePort,
		Tags:    []string{"api", "user"},
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/api/v1/user-service/health", serviceAddress, servicePort),
			Method:                         "GET",
			Interval:                       "10s",
			Timeout:                        "5s",
			Notes:                          "Basic health check in user service " + fmt.Sprintf("http://%s:%d/api/v1/health", serviceAddress, servicePort),
			DeregisterCriticalServiceAfter: "1m",
		},
	}

	global.Logger.Info("Registering service with Consul", zap.Any("registration", registration))

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		// global.Logger.Error("Failed to connect to Consul:", zap.Error(err))
		panic(err)
	}
	global.ConsulClient = consulClient

	err = global.ConsulClient.Agent().ServiceRegister(registration)
	if err != nil {
		global.Logger.Error("Failed to register service:", zap.Error(err))
		panic(err)
	}

	return func() {
		if err := global.ConsulClient.Agent().ServiceDeregister(serviceID); err != nil {
			global.Logger.Error("Failed to deregister service", zap.Error(err))
		} else {
			global.Logger.Info("Service deregistered")
		}
	}
}
