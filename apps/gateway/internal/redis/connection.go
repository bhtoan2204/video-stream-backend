package redis

import (
	"context"
	"fmt"

	"github.com/bhtoan2204/gateway/global"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	rc := global.Config.RedisConfig

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", rc.Host, rc.Port),
		Password: rc.Password, // no password set
		DB:       rc.Database, // use default DB
		PoolSize: 10,
	})

	_, err := rdb.Ping(global.Ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis:", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("Connected to Redis successfully")
	global.Redis = rdb
}

func CacheData(ctx context.Context, key string, value string, ttl int) {
	tracer := otel.Tracer("gateway/redis")
	ctx, span := tracer.Start(ctx, "Redis SET")
	span.SetAttributes(
		attribute.String("cache.key", key),
		attribute.String("cache.value", value),
		attribute.Int("cache.ttl_seconds", ttl),
	)
	defer span.End()
	err := global.Redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		span.RecordError(err)
		global.Logger.Error("Failed to cache data in Redis", zap.Error(err))
	}
}

func GetData(ctx context.Context, key string) (string, error) {
	tracer := otel.Tracer("gateway/redis")
	ctx, span := tracer.Start(ctx, "Redis GET")
	span.SetAttributes(attribute.String("cache.key", key))
	defer span.End()
	data, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		span.RecordError(err)
		global.Logger.Error("Failed to get data from Redis", zap.Error(err))
		return "", err
	}
	return data, nil
}
