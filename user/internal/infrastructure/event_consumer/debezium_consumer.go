package consumer

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/event"
	publishedEvent "github.com/bhtoan2204/user/internal/application/event/event"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/persistent_object"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type DebeziumMessage struct {
	Payload struct {
		After  json.RawMessage `json:"after"`
		Op     string          `json:"op"`
		Source struct {
			Name string `json:"name"`
		} `json:"source"`
	} `json:"payload"`
}

type DebeziumConsumer struct {
	readers  []*kafka.Reader
	eventBus *event.EventBus
}

func NewDebezium(eventBus *event.EventBus) *DebeziumConsumer {
	topics := []string{
		"dbserver1.user." + persistent_object.ActivityLog{}.TableName(),
		"dbserver1.user." + persistent_object.Permission{}.TableName(),
		"dbserver1.user." + persistent_object.RefreshToken{}.TableName(),
		"dbserver1.user." + persistent_object.Role{}.TableName(),
		"dbserver1.user." + persistent_object.UserSettings{}.TableName(),
		"dbserver1.user." + persistent_object.User{}.TableName(),
	}

	var readers []*kafka.Reader

	for _, topic := range topics {
		global.Logger.Info("Creating reader", zap.String("topic", topic))
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{global.Config.KafkaConfig.Broker},
			// Brokers: []string{"kafka:29092"},
			GroupID: "mysql-user-connector",
			Topic:   topic,
		})
		readers = append(readers, reader)
	}

	return &DebeziumConsumer{
		readers:  readers,
		eventBus: eventBus,
	}
}

func (d *DebeziumConsumer) ProcessMessage(msg kafka.Message, topic string) {
	var message DebeziumMessage
	if err := json.Unmarshal(msg.Value, &message); err != nil {
		global.Logger.Error("Error unmarshalling message", zap.Error(err))
		return
	}

	tableName := getTableName(topic)

	if tableName == "users" && (message.Payload.Op == "c" || message.Payload.Op == "u") {
		var ru struct {
			ID           string `json:"id"`
			CreatedAt    int64  `json:"created_at"`
			UpdatedAt    int64  `json:"updated_at"`
			DeletedAt    *int64 `json:"deleted_at"`
			Username     string `json:"username"`
			Email        string `json:"email"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Phone        string `json:"phone"`
			BirthDate    int64  `json:"birth_date"`
			Address      string `json:"address"`
			PasswordHash string `json:"password_hash"`
			PinCode      string `json:"pin_code"`
			Status       int    `json:"status"`
		}
		if err := json.Unmarshal(message.Payload.After, &ru); err != nil {
			global.Logger.Error("Error unmarshalling user", zap.Error(err))
			return
		}

		indexUserEvent := &publishedEvent.IndexUserEvent{
			ID:           ru.ID,
			CreatedAt:    time.UnixMilli(ru.CreatedAt),
			UpdatedAt:    time.UnixMilli(ru.UpdatedAt),
			DeletedAt:    int64ToTimePtr(ru.DeletedAt),
			Username:     ru.Username,
			Email:        ru.Email,
			FirstName:    ru.FirstName,
			LastName:     ru.LastName,
			Phone:        ru.Phone,
			BirthDate:    int64ToTimePtr(&ru.BirthDate),
			Address:      ru.Address,
			PasswordHash: ru.PasswordHash,
			PinCode:      ru.PinCode,
			Status:       1,
		}

		if _, err := d.eventBus.Dispatch(indexUserEvent); err != nil {
			global.Logger.Error("Error dispatching event", zap.Error(err))
			return
		}
		global.Logger.Info("Event dispatched", zap.Any("event", indexUserEvent))
	}
}

func (d *DebeziumConsumer) Consume() {
	for _, reader := range d.readers {
		go func(r *kafka.Reader) {
			for {
				m, err := r.ReadMessage(context.Background())
				if err != nil {
					global.Logger.Error("Error reading message", zap.String("topic", r.Config().Topic), zap.Error(err))
					continue
				}
				global.Logger.Info("Message received", zap.String("topic", r.Config().Topic))
				d.ProcessMessage(m, r.Config().Topic)
			}
		}(reader)
	}
}

func int64ToTimePtr(timestamp *int64) *time.Time {
	if timestamp == nil {
		return nil
	}

	t := time.UnixMilli(*timestamp)
	return &t
}

func getTableName(fullName string) string {
	parts := strings.Split(fullName, ".")
	if len(parts) > 2 {
		return parts[len(parts)-1]
	}
	return fullName
}
