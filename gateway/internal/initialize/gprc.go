package initialize

import (
	"github.com/bhtoan2204/gateway/global"
	"google.golang.org/grpc"
)

func InitGrpcServer() {
	s := grpc.NewServer()
	global.GrpcServer = s
}
