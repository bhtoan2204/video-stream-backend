package initialize

import (
	"net"
	"os"

	"github.com/bhtoan2204/video/global"
	"go.uber.org/zap"
)

func InitListener() {
	listener, err := net.Listen("tcp", ":1111")
	if err != nil {
		global.Logger.Error("Failed to allocate port", zap.Error(err))
		os.Exit(1)
	}
	global.Listener = listener
}
