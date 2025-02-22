package consumer

import (
	"context"
	"encoding/json"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/domain/event"
	"github.com/bhtoan2204/user/internal/domain/repository/shared"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func ConsumeMessage(esRepositories *shared.Repositories) {
	for {
		msg, err := global.KafkaConsumer.ReadMessage(context.Background())
		if err != nil {
			global.Logger.Error("Failed to consume message from Kafka", zap.Error(err))
			continue
		}
		global.Logger.Info("Message consumed from Kafka", zap.String("key", string(msg.Key)), zap.String("value", string(msg.Value)))
		processMessage(msg, esRepositories)
	}
}

func processMessage(msg kafka.Message, esRepositories *shared.Repositories) {
	switch string(msg.Key) {
	case "UserCreated":
		var user entities.User
		var userCreatedEvent event.UserCreatedEvent
		if err := json.Unmarshal(msg.Value, &userCreatedEvent); err != nil {
			global.Logger.Error("Failed to unmarshal UserCreatedEvent", zap.Error(err))
			return
		}
		user.ID = userCreatedEvent.Payload.ID
		user.Username = userCreatedEvent.Payload.Username
		user.Email = userCreatedEvent.Payload.Email
		user.FirstName = userCreatedEvent.Payload.FirstName
		user.LastName = userCreatedEvent.Payload.LastName
		if err := esRepositories.ESUserRepository.Index(&user); err != nil {
			global.Logger.Error("Failed to index user", zap.Error(err))
		}
	default:
		global.Logger.Warn("Unknown message key", zap.String("key", string(msg.Key)))
	}
}
