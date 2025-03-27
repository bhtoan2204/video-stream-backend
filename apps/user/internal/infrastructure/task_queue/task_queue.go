package task_queue

import (
	"fmt"

	"github.com/bhtoan2204/user/global"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func InitAsynq() func() {
	redisConfig := global.Config.RedisConfig
	redisAddr := fmt.Sprintf("%s:%v", redisConfig.Host, redisConfig.Port)
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: redisAddr,
	})
	global.AsynqClient = client
	global.Logger.Info("Connected to Asynq successfully")

	return func() {
		if err := client.Close(); err != nil {
			global.Logger.Error("Failed to close Asynq client", zap.Error(err))
		}
	}
}
