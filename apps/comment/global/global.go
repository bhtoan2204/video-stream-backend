package global

import (
	"net"

	"github.com/bhtoan2204/comment/internal/infrastructure/grpc/proto/video"
	"github.com/bhtoan2204/comment/pkg/logger"
	"github.com/bhtoan2204/comment/pkg/settings"
	"github.com/gocql/gocql"
	"github.com/hashicorp/consul/api"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"
)

var (
	Listener        net.Listener
	Config          settings.Config
	Logger          *logger.LoggerZap
	Redis           *redis.Client
	ConsulClient    *api.Client
	MDB             *gorm.DB
	ScyllaSession   *gocql.Session
	KafkaProducer   *kafka.Writer
	KafkaConsumer   *kafka.Reader
	VideoGRPCClient video.VideoServiceClient
)
