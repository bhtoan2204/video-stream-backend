package handler

import (
	"github.com/bhtoan2204/user/internal/application/command/command"
	common "github.com/bhtoan2204/user/internal/application/common/command"
	"github.com/bhtoan2204/user/internal/application/interfaces"
)

type RefreshTokenCommandHandler struct {
	userService interfaces.UserServiceInterface
}

func NewRefreshTokenCommandHandler(userService interfaces.UserServiceInterface) *RefreshTokenCommandHandler {
	return &RefreshTokenCommandHandler{
		userService: userService,
	}
}

func (h *RefreshTokenCommandHandler) Handle(cmd *command.RefreshTokenCommand) (*common.RefreshTokenCommandResult, error) {
	return h.userService.Refresh(cmd)
}
