package event

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/domain/event"
	"github.com/segmentio/kafka-go"
)

// KafkaDomainEventPublisher implement DomainEventPublisher sử dụng Kafka.
type KafkaDomainEventPublisher struct {
	writer *kafka.Writer
}

// NewKafkaDomainEventPublisher khởi tạo KafkaDomainEventPublisher với danh sách broker và topic.
func NewKafkaDomainEventPublisher() *KafkaDomainEventPublisher {
	kafkaConfig := global.Config.KafkaConfig
	brokers := []string{kafkaConfig.Broker}
	topic := kafkaConfig.Topic

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &KafkaDomainEventPublisher{writer: writer}
}

// Publish chuyển domain event thành JSON và gửi message lên Kafka.
func (p *KafkaDomainEventPublisher) Publish(event event.DomainEvent) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	msg := kafka.Message{
		Key:   []byte(event.EventType()),
		Value: payload,
		Time:  time.Now(),
	}

	return p.writer.WriteMessages(context.Background(), msg)
}
