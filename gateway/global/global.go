package global

import (
	"github.com/bhtoan2204/gateway/pkg/logger"
	"github.com/bhtoan2204/gateway/pkg/settings"
	"github.com/hashicorp/consul/api"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"

	"github.com/redis/go-redis/v9"
)

var (
	Config        settings.Config
	Logger        *logger.LoggerZap
	Redis         *redis.Client
	ConsulClient  *api.Client
	KafkaProducer *kafka.Writer
	KafkaConsumer *kafka.Reader
	GrpcServer    *grpc.Server
	// S3Client *s3.Client
)
