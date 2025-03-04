package handler

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/command/command"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
)

type LoginCommandHandler struct {
	userService interfaces.UserServiceInterface
}

func NewLoginCommandHandler(userService interfaces.UserServiceInterface) *LoginCommandHandler {
	return &LoginCommandHandler{
		userService: userService,
	}
}

func (h *LoginCommandHandler) Handle(ctx context.Context, cmd *command.LoginCommand) (*command.LoginCommandResult, error) {
	return h.userService.Login(ctx, cmd)
}
