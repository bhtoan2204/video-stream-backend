package initialize

import (
	"os"
	"user/global"
)

func Run() {
	InitListener()
	InitConsul()
	r := InitRouter()
	// listener, err := net.Listen("tcp", ":0")
	// if err != nil {
	// 	// global.Logger.Error("Failed to allocate port", zap.Error(err))
	// 	os.Exit(1)
	// }
	// port := listener.Addr().(*net.TCPAddr).Port
	// global.Logger.Info("Allocated port", zap.Int("port", port))
	// fmt.Println("Allocated port", port)
	if err := r.RunListener(global.Listener); err != nil {
		// global.Logger.Error("Failed to start server", zap.Error(err))
		// Handle error
		os.Exit(1)
	}
}
