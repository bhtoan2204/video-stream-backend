package initialize

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bhtoan2204/user/global"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type DebeziumMessage struct {
	Source struct {
		Database string `json:"db"`
		Table    string `json:"table"`
	} `json:"source"`
	Payload struct {
		Before json.RawMessage `json:"before"`
		After  json.RawMessage `json:"after"`
		Op     string          `json:"op"` // "c", "u", "d"
	} `json:"payload"`
}

func processDebeziumMessage(msg kafka.Message) {
	global.Logger.Info("Processing message", zap.String("key", string(msg.Key)))

	var message DebeziumMessage

	if err := json.Unmarshal(msg.Value, &message); err != nil {
		fmt.Println("Error unmarshalling message", err)
		return
	}

	switch message.Payload.Op {
	case "c", "u":
		global.Logger.Info("Create or Update operation")
		var data map[string]interface{}
		if err := json.Unmarshal(message.Payload.After, &data); err != nil {
			global.Logger.Error("Error unmarshalling user data", zap.Error(err))
			return
		}

		docID := fmt.Sprintf("%v", data["id"])
		res, err := global.ESClient.Index(
			message.Source.Table, // index name in ES
			strings.NewReader(string(message.Payload.After)),
			global.ESClient.Index.WithDocumentID(docID),
			global.ESClient.Index.WithContext(context.Background()),
		)
		if err != nil {
			global.Logger.Error("Error indexing document", zap.Error(err))
			return
		}
		res.Body.Close()
		global.Logger.Info("Document indexed successfully")
	default:
		global.Logger.Info("No operation")
	}
}

func ConsumeCDC() {
	for {
		fmt.Println("Waiting for message")
		m, err := global.KafkaDebeziumConsumer.ReadMessage(context.Background())
		fmt.Println("Message received", string(m.Value))
		if err != nil {
			fmt.Println("Error reading message", err)
			continue
		}

		processDebeziumMessage(m)
	}
}

// name: mysql-connector
func InitDebeziumConsumer() {
	global.KafkaDebeziumConsumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{global.Config.KafkaConfig.Broker},
		GroupID: "inventory-connector",
		Topic:   "dbserver1.user.users",
	})

	go ConsumeCDC()
}
