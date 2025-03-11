package interfaces

import (
	"context"

	common "github.com/bhtoan2204/user/internal/application/common/event"
	"github.com/bhtoan2204/user/internal/application/event_bus/event"
)

type UserListenerInterface interface {
	IndexUser(ctx context.Context, indexUserEvent *event.IndexUserEvent) (*common.IndexResult, error)
}
