package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bhtoan2204/worker/internal/payload"
	"github.com/hibiken/asynq"
)

func HandleImageResizeTask(ctx context.Context, t *asynq.Task) error {
	var payload payload.ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("error decrypting payload: %v", err)
	}

	fmt.Printf("processing image: %s\n", payload.SourcePath)
	return nil
}
