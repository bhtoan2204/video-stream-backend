package handler

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/command/command"
	common "github.com/bhtoan2204/user/internal/application/common/command"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
)

type LogoutCommandHandler struct {
	userService interfaces.UserServiceInterface
}

func NewLogoutCommandHandler(userService interfaces.UserServiceInterface) *LogoutCommandHandler {
	return &LogoutCommandHandler{
		userService: userService,
	}
}

func (h *LogoutCommandHandler) Handle(ctx context.Context, cmd *command.LogoutCommand) (*common.LogoutCommandResult, error) {
	return h.userService.Logout(ctx, cmd)
}
