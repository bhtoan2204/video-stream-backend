package initialize

import (
	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LogConfig)
	global.Logger.Info("Logger initialized successfully")
}
