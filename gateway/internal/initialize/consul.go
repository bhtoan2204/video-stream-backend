package initialize

import (
	"github.com/bhtoan2204/gateway/global"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

func InitConsul() {
	globalConsulConfig := global.Config.ConsulConfig

	consulConfig := api.DefaultConfig()
	consulConfig.Address = globalConsulConfig.Address
	consulConfig.Scheme = globalConsulConfig.Scheme
	consulConfig.Datacenter = globalConsulConfig.DataCenter
	consulConfig.Token = globalConsulConfig.Token

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		global.Logger.Error("Failed to connect to Consul:", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("Connected to Consul successfully")
	global.ConsulClient = consulClient
}
