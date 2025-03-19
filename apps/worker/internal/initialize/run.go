package initialize

import (
	"log"

	"github.com/bhtoan2204/worker/internal/handler"
	"github.com/bhtoan2204/worker/internal/tasks"
	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

func Run() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
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

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
