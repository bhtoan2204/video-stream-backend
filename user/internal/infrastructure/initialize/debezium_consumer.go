package initialize

import (
	"github.com/bhtoan2204/user/global"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

// name: mysql-connector
func InitDebeziumConsumer() {
	global.KafkaDebeziumConsumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{global.Config.KafkaConfig.Broker},
		GroupID: "mysql_server.appdb.user",
		Topic:   "debezium-consumer-group",
	})

	global.Logger.Info("Debezium consumer initialized", zap.String("broker", global.Config.KafkaConfig.Broker), zap.String("group_id", "mysql_server.appdb.user"), zap.String("topic", "debezium-consumer-group"))
}
