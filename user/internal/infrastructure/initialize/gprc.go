package initialize

import (
	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/infrastructure/grpc/proto/user"
	serviceserver "github.com/bhtoan2204/user/internal/infrastructure/service_server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitGrpcServer() {
	// listener, err := net.Listen("tcp", ":0")
	// if err != nil {
	// 	global.Logger.Fatal("Failed to allocate port", zap.Error(err))
	// }

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, serviceserver.NewUserServiceServer())

	go func() {
		if err := grpcServer.Serve(global.Listener); err != nil {
			global.Logger.Error("Error serving gRPC server", zap.Error(err))
		}
	}()
}
