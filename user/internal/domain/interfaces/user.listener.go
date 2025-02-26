package interfaces

import (
	common "github.com/bhtoan2204/user/internal/application/common/event"
	"github.com/bhtoan2204/user/internal/application/event/event"
)

type UserListenerInterface interface {
	IndexUser(indexUserEvent *event.IndexUserEvent) (*common.IndexResult, error)
}
