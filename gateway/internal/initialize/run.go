package initialize

import (
	"os"
	"strconv"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/internal/redis"

	"go.uber.org/zap"
)

func Run() {
	InitLoadConfig()
	InitLogger()
	InitConsul()
	redis.InitRedis()
	// InitKafka()
	InitUserGRPCClient()
	tracerShutdown := InitProvider()
	defer tracerShutdown()

	r := InitRouter()

	global.Logger.Info("Initialize all services successfully")
	if err := r.Run(":" + strconv.Itoa(global.Config.Server.Port)); err != nil {
		global.Logger.Error("Failed to start server", zap.Error(err))
		os.Exit(1)
	}
}
