package initialize

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/domain/repository/query"
	"github.com/bhtoan2204/user/internal/infrastructure/db/elasticsearch/repository"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type DebeziumMessage struct {
	Payload struct {
		Before json.RawMessage `json:"before"`
		After  json.RawMessage `json:"after"`
		Source struct {
			Version   string `json:"version"`
			Connector string `json:"connector"`
			Name      string `json:"name"`
		} `json:"source"`
		Transaction json.RawMessage `json:"transaction"`
		Op          string          `json:"op"`
		TsMs        int64           `json:"ts_ms"`
		TsUs        int64           `json:"ts_us"`
		TsNs        int64           `json:"ts_ns"`
	} `json:"payload"`
}

type rawUser struct {
	ID           uint   `json:"id"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	DeletedAt    *int64 `json:"deleted_at"` // pointer để nhận null
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

type Debezium struct {
	reader           *kafka.Reader
	esUserRepository query.ESUserRepository
}

func NewDebezium(esRepo query.ESUserRepository) *Debezium {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{global.Config.KafkaConfig.Broker},
		GroupID: "inventory-connector",
		Topic:   "dbserver1.user.users",
	})
	return &Debezium{
		reader:           reader,
		esUserRepository: esRepo,
	}
}

func (d *Debezium) ProcessMessage(msg kafka.Message) {
	global.Logger.Info("Processing message")
	var message DebeziumMessage
	if err := json.Unmarshal(msg.Value, &message); err != nil {
		global.Logger.Error("Error unmarshalling message", zap.Error(err))
		return
	}
	fmt.Println(string(message.Payload.After))
	switch message.Payload.Op {
	case "c", "u":
		var ru rawUser
		if err := json.Unmarshal(message.Payload.After, &ru); err != nil {
			global.Logger.Error("Error unmarshalling raw user data", zap.Error(err))
			return
		}

		var user entities.User
		user.AbstractModel = entities.AbstractModel{
			ID: uint(ru.ID),
			// CreatedAt: ru.CreatedAt,
			// UpdatedAt: ru.UpdatedAt,
		}
		user.Username = ru.Username
		user.Email = ru.Email
		user.FirstName = ru.FirstName
		user.LastName = ru.LastName
		user.Phone = ru.Phone
		user.Address = ru.Address
		user.PasswordHash = ru.PasswordHash
		user.PinCode = ru.PinCode
		user.Status = 1
		user.BirthDate = nil
		// if ru.BirthDate > 0 {
		// 	t := time.Unix(0, ru.BirthDate*int64(time.Millisecond))
		// 	user.BirthDate = &t
		// } else {
		// 	user.BirthDate = nil
		// }

		if err := d.esUserRepository.Index(&user); err != nil {
			global.Logger.Error("Error indexing document", zap.Error(err))
			return
		}
		global.Logger.Info("Document indexed successfully")
	default:
		global.Logger.Info("No operation")
	}
}

func (d *Debezium) Consume() {
	for {
		global.Logger.Info("Waiting for message")
		m, err := d.reader.ReadMessage(context.Background())
		if err != nil {
			global.Logger.Error("Error reading message", zap.Error(err))
			continue
		}
		global.Logger.Info("Message received")
		d.ProcessMessage(m)
	}
}

func InitDebeziumConsumer() {

	esRepo := repository.NewESUserRepository(global.ESClient)
	debezium := NewDebezium(esRepo)
	go debezium.Consume()
}
