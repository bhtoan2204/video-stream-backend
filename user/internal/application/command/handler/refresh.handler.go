package handler

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/command/command"
	common "github.com/bhtoan2204/user/internal/application/common/command"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
)

type RefreshTokenCommandHandler struct {
	userService interfaces.UserServiceInterface
}

func NewRefreshTokenCommandHandler(userService interfaces.UserServiceInterface) *RefreshTokenCommandHandler {
	return &RefreshTokenCommandHandler{
		userService: userService,
	}
}

func (h *RefreshTokenCommandHandler) Handle(ctx context.Context, cmd *command.RefreshTokenCommand) (*common.RefreshTokenCommandResult, error) {
	return h.userService.Refresh(ctx, cmd)
}
