package global

import (
	"net"

	"github.com/bhtoan2204/user/pkg/logger"
	"github.com/bhtoan2204/user/pkg/settings"
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"
)

var (
	Listener     net.Listener
	Config       settings.Config
	Logger       *logger.LoggerZap
	Redis        *redis.Client
	ConsulClient *api.Client
	MDB          *gorm.DB
	// KafkaProducer *kafka.Writer
	// KafkaConsumer *kafka.Reader
	// S3Client *s3.Client
)
