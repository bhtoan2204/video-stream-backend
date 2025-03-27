package initialize

import (
	"github.com/bhtoan2204/video/global"
	"github.com/bhtoan2204/video/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LogConfig)
	global.Logger.Info("Logger initialized successfully")
}
