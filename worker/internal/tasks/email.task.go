package tasks

import (
	"encoding/json"

	"github.com/bhtoan2204/worker/internal/payload"
	"github.com/hibiken/asynq"
)

const TypeEmailDelivery = "email:deliver"

func NewEmailTask(payload payload.EmailPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, data), nil
}
