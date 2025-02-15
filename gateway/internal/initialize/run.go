package initialize

import (
	"os"
	"strconv"

	"github.com/bhtoan2204/gateway/global"

	"go.uber.org/zap"
)

func Run() {
	InitLogger()
	InitLoadConfig()
	InitConsul()
	InitRedis()
	r := InitRouter()
	global.Logger.Info("Initialize all services successfully")
	if err := r.Run(":" + strconv.Itoa(global.Config.Server.Port)); err != nil {
		global.Logger.Error("Failed to start server", zap.Error(err))
		// Handle error
		os.Exit(1)
	}
}
