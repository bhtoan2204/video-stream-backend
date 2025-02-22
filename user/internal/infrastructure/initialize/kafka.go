package initialize

import (
	"context"

	"github.com/bhtoan2204/user/global"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaMessage struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func InitKafkaProducer() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP(global.Config.KafkaConfig.Broker),
		Topic:    global.Config.KafkaConfig.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	global.Logger.Info("Kafka producer initialized", zap.String("broker", global.Config.KafkaConfig.Broker), zap.String("topic", global.Config.KafkaConfig.Topic))
}

func InitKafkaConsumer() {
	global.KafkaConsumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{global.Config.KafkaConfig.Broker},
		GroupID: global.Config.KafkaConfig.GroupID,
		Topic:   global.Config.KafkaConfig.Topic,
	})

	global.Logger.Info("Kafka consumer initialized", zap.String("broker", global.Config.KafkaConfig.Broker), zap.String("group_id", global.Config.KafkaConfig.GroupID), zap.String("topic", global.Config.KafkaConfig.Topic))

}

func ProduceMessage(key, message string) error {
	err := global.KafkaProducer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(message),
		},
	)
	if err != nil {
		global.Logger.Error("Failed to produce message to Kafka", zap.Error(err))
		return err
	}
	global.Logger.Info("Message sent to Kafka", zap.String("key", key), zap.String("value", message))
	return nil
}

func InitKafka() {
	InitKafkaProducer()
	InitKafkaConsumer()
	global.Logger.Info("Kafka initialized", zap.Any("KafkaConfig", global.Config.KafkaConfig))
}
