package global

import (
	"net"

	"github.com/bhtoan2204/video/pkg/logger"
	"github.com/bhtoan2204/video/pkg/settings"
	"github.com/bhtoan2204/video/third_party"
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
	S3Client     *third_party.S3Client
	// KafkaProducer *kafka.Writer
	// KafkaConsumer *kafka.Reader
)
