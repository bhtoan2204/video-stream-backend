package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/bhtoan2204/worker/internal/payload"
	"github.com/bhtoan2204/worker/pkg/email"
	"github.com/hibiken/asynq"
)

func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var payload payload.EmailPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("error decrypting payload: %v", err)
	}

	if err := email.SendEmail(payload); err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	log.Printf("sending successful to: %s", payload.To)
	return nil
}
