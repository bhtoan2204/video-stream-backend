package initialize

import (
	"context"
	"fmt"

	"github.com/bhtoan2204/gateway/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	rc := global.Config.RedisConfig
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", rc.Host, rc.Port),
		Password: rc.Password, // no password set
		DB:       rc.Database, // use default DB
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis:", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("Connected to Redis successfully")
	global.Redis = rdb
}
