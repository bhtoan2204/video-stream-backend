package handler

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/command_bus/command"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
)

type UpdateUserCommandHandler struct {
	userService interfaces.UserServiceInterface
}

func NewUpdateUserCommandHandler(userService interfaces.UserServiceInterface) *UpdateUserCommandHandler {
	return &UpdateUserCommandHandler{userService: userService}
}

func (h *UpdateUserCommandHandler) Handle(ctx context.Context, cmd *command.UpdateUserCommand) (*command.UpdateUserCommandResult, error) {
	return h.userService.UpdateUser(ctx, cmd)
}
