package tasks

import (
	"encoding/json"

	"github.com/bhtoan2204/worker/internal/payload"
	"github.com/hibiken/asynq"
)

const TypeVideoTranscoding = "video:transcoding"

func NewVideoTranscodingTask(payload payload.VideoTranscodingPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeVideoTranscoding, data), nil
}
