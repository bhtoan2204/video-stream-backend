package initialize

import (
	"net"
	"os"

	"github.com/bhtoan2204/user/global"
)

func InitListener() {
	listener, err := net.Listen("tcp", "localhost:8090")
	// fmt.Println("Listening on", listener.Addr().String())
	if err != nil {
		// global.Logger.Error("Failed to allocate port", zap.Error(err))
		os.Exit(1)
	}
	global.Listener = listener
}
