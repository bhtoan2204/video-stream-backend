package global

import (
	"context"

	"github.com/bhtoan2204/gateway/pkg/grpc/proto/user"
	"github.com/bhtoan2204/gateway/pkg/logger"
	"github.com/bhtoan2204/gateway/pkg/settings"
	"github.com/hashicorp/consul/api"
	"github.com/segmentio/kafka-go"

	"github.com/redis/go-redis/v9"
)

var (
	Config         settings.Config
	Logger         *logger.LoggerZap
	Redis          *redis.Client
	ConsulClient   *api.Client
	KafkaProducer  *kafka.Writer
	KafkaConsumer  *kafka.Reader
	UserGRPCClient user.UserServiceClient
	Ctx            = context.Background()
)
