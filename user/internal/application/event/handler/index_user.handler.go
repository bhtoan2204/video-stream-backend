package handler

import (
	"context"

	common "github.com/bhtoan2204/user/internal/application/common/event"
	"github.com/bhtoan2204/user/internal/application/event/event"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
)

type IndexUserEventHandler struct {
	userListener interfaces.UserListenerInterface
}

func NewIndexUserEventHandler(userListener interfaces.UserListenerInterface) *IndexUserEventHandler {
	return &IndexUserEventHandler{userListener: userListener}
}

func (h *IndexUserEventHandler) Handle(ctx context.Context, event *event.IndexUserEvent) (*common.IndexResult, error) {
	return h.userListener.IndexUser(ctx, event)
}
