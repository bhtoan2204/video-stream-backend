package initialize

import (
	"github.com/bhtoan2204/comment/global"
	"github.com/bhtoan2204/comment/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LogConfig)
	global.Logger.Info("Logger initialized successfully")
}
