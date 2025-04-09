package initialize

import (
	"log"
	"strconv"

	"github.com/bhtoan2204/worker/global"
	"github.com/bhtoan2204/worker/internal/handler"
	loadconfig "github.com/bhtoan2204/worker/internal/loadConfig"
	"github.com/bhtoan2204/worker/internal/logger"
	"github.com/bhtoan2204/worker/internal/storage"
	"github.com/bhtoan2204/worker/internal/tasks"
	"github.com/bhtoan2204/worker/internal/tracing"
	"github.com/hibiken/asynq"
)

func Run() {
	loadconfig.InitLoadConfig()
	logger.InitLogger()

	tracerShutdown := tracing.InitProvider()
	defer tracerShutdown()

	storage.InitStorageService()

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: global.Config.RedisConfig.Host + ":" + strconv.Itoa(global.Config.RedisConfig.Port)},
		asynq.Config{
			Concurrency: 20,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailDelivery, handler.HandleEmailDeliveryTask)
	mux.HandleFunc(tasks.TypeImageResize, handler.HandleImageResizeTask)
	mux.HandleFunc(tasks.TypeVideoTranscoding, handler.HandleVideoTranscodingTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
