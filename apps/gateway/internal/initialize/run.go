package initialize

import (
	"os"
	"strconv"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/internal/consul"
	"github.com/bhtoan2204/gateway/internal/redis"
	"github.com/bhtoan2204/gateway/internal/router"

	"go.uber.org/zap"
)

func Run() {
	InitLoadConfig()
	InitLogger()
	consul.InitConsul()
	redis.InitRedis()
	// InitKafka()
	InitUserGRPCClient()
	tracerShutdown := InitProvider()
	defer tracerShutdown()

	r := router.InitRouter()

	global.Logger.Info("Initialize all services successfully")
	if err := r.Run(":" + strconv.Itoa(global.Config.Server.Port)); err != nil {
		global.Logger.Error("Failed to start server", zap.Error(err))
		os.Exit(1)
	}
}
