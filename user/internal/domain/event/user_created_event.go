package event

import (
	"time"

	common "github.com/bhtoan2204/user/internal/application/common/command"
)

type UserCreatedEvent struct {
	Payload  *common.UserResult `json:"payload"`
	Occurred time.Time          `json:"occurred"`
}

func (e UserCreatedEvent) EventType() string {
	return "UserCreated"
}

func (e UserCreatedEvent) OccurredAt() time.Time {
	return e.Occurred
}
