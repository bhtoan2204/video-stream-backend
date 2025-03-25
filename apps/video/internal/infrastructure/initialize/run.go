package initialize

import (
	"os"

	"github.com/bhtoan2204/video/global"
	"github.com/bhtoan2204/video/internal/infrastructure/db/mysql"
	"github.com/bhtoan2204/video/internal/infrastructure/db/mysql/repository"
	"github.com/bhtoan2204/video/internal/infrastructure/db/scylla"
	"github.com/bhtoan2204/video/internal/infrastructure/grpc"
	"github.com/bhtoan2204/video/internal/infrastructure/redis"
	"github.com/bhtoan2204/video/internal/infrastructure/tracing"
)

func Run() {
	InitLoadConfig()
	InitLogger()
	InitListener()

	redis.InitRedis()

	deregisterService := InitConsul()
	defer deregisterService()

	mysql.InitDB()
	defer mysql.CloseDB()

	scylla.InitDB()
	defer scylla.CloseDB()

	InitStorageService()

	tracerShutdown := tracing.InitProvider()
	defer tracerShutdown()

	grpc.StartGrpcServer(repository.NewVideoRepository(global.MDB))

	r := InitRouter()
	if err := r.RunListener(global.Listener); err != nil {
		os.Exit(1)
	}
}
