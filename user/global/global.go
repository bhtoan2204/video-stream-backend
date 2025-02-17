package global

import (
	"net"
	"user/pkg/settings"

	"github.com/hashicorp/consul/api"

	"github.com/redis/go-redis/v9"
)

var (
	Listener net.Listener
	Config   settings.Config
	// Logger       *logger.LoggerZap
	Redis        *redis.Client
	ConsulClient *api.Client
	// KafkaProducer *kafka.Writer
	// KafkaConsumer *kafka.Reader
	// S3Client *s3.Client
)
