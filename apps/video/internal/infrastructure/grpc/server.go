package grpc

import (
	"fmt"
	"net"

	"github.com/bhtoan2204/video/global"
	repository "github.com/bhtoan2204/video/internal/domain/repository/command"
	"github.com/bhtoan2204/video/internal/infrastructure/grpc/proto/video"
	service_server "github.com/bhtoan2204/video/internal/infrastructure/grpc/service_server"
	"github.com/bhtoan2204/video/utils"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func StartGrpcServer(videoRepository repository.VideoRepositoryInterface) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		global.Logger.Fatal("Failed to allocate port", zap.Error(err))
	}

	grpcServer := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))

	video.RegisterVideoServiceServer(grpcServer, service_server.NewVideoServiceServer(videoRepository))

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	healthServer.SetServingStatus("video.VideoService", grpc_health_v1.HealthCheckResponse_SERVING)

	globalConsulConfig := global.Config.ConsulConfig

	consulConfig := api.DefaultConfig()
	consulConfig.Address = globalConsulConfig.Address
	consulConfig.Scheme = globalConsulConfig.Scheme
	consulConfig.Datacenter = globalConsulConfig.DataCenter
	consulConfig.Token = globalConsulConfig.Token

	serviceID := uuid.New().String()
	servicePort := listener.Addr().(*net.TCPAddr).Port
	serviceAddress, err := utils.GetInternalIP()
	if err != nil {
		global.Logger.Error("Failed to get internal IP address:", zap.Error(err))
		panic(err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "video-grpc",
		Address: consulConfig.Address,
		Port:    servicePort,
		Tags:    []string{"grpc", "video"},
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/api/v1/video-service/health", serviceAddress, global.Listener.Addr().(*net.TCPAddr).Port),
			Method:                         "GET",
			Interval:                       "10s",
			Timeout:                        "5s",
			Notes:                          "Basic health check in video grpc " + fmt.Sprintf("http://%s:%d/api/v1/health", serviceAddress, servicePort),
			DeregisterCriticalServiceAfter: "1m",
		},
	}

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		// global.Logger.Error("Failed to connect to Consul:", zap.Error(err))
		panic(err)
	}

	global.GrpcConsulClient = consulClient
	err = global.GrpcConsulClient.Agent().ServiceRegister(registration)

	if err != nil {
		global.Logger.Error("Failed to register service:", zap.Error(err))
		panic(err)
	}

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			global.Logger.Error("Error serving gRPC server", zap.Error(err))
		}
	}()
}
