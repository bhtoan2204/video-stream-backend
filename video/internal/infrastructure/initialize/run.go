package initialize

import (
	"context"
	"os"

	"github.com/bhtoan2204/video/global"
	s3client "github.com/bhtoan2204/video/third_party"
)

func Run() {
	InitListener()
	// InitConsul()
	r := InitRouter()
	// listener, err := net.Listen("tcp", ":0")
	// if err != nil {
	// 	// global.Logger.Error("Failed to allocate port", zap.Error(err))
	// 	os.Exit(1)
	// }
	// port := listener.Addr().(*net.TCPAddr).Port
	// global.Logger.Info("Allocated port", zap.Int("port", port))
	// fmt.Println("Allocated port", port)
	InitStorageService()
	s3client.NewS3Client(context.Background(), global.Config.S3Config.Bucket, global.Config.S3Config.Region)
	if err := r.RunListener(global.Listener); err != nil {
		// global.Logger.Error("Failed to start server", zap.Error(err))
		// Handle error
		os.Exit(1)
	}
}
