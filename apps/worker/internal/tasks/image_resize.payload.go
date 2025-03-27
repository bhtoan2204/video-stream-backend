package tasks

import (
	"encoding/json"

	"github.com/bhtoan2204/worker/internal/payload"
	"github.com/hibiken/asynq"
)

const TypeImageResize = "image:resize"

func NewImageResizeTask(payload payload.ImageResizePayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeImageResize, data), nil
}
