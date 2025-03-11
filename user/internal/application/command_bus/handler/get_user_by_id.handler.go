package handler

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/command_bus/command"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
)

type GetUserByIdCommandHandler struct {
	userService interfaces.UserServiceInterface
}

func NewGetUserByIdCommandHandler(userService interfaces.UserServiceInterface) *GetUserByIdCommandHandler {
	return &GetUserByIdCommandHandler{
		userService: userService,
	}
}

func (h *GetUserByIdCommandHandler) Handle(ctx context.Context, cmd *command.GetUserByIdCommand) (*command.GetUserByIdCommandResult, error) {
	return h.userService.GetUserById(ctx, cmd)
}
