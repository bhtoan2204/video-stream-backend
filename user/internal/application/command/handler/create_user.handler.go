package handler

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/command/command"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
)

type CreateUserCommandHandler struct {
	userService interfaces.UserServiceInterface
}

func NewCreateUserCommandHandler(userService interfaces.UserServiceInterface) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		userService: userService,
	}
}

func (h *CreateUserCommandHandler) Handle(ctx context.Context, cmd *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	return h.userService.CreateUser(ctx, cmd)
}
