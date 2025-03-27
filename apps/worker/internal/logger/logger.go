package logger

import (
	"github.com/bhtoan2204/worker/global"
	"github.com/bhtoan2204/worker/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LogConfig)
	global.Logger.Info("Logger initialized successfully")
}
