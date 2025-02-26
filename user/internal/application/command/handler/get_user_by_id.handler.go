package handler

import (
	"github.com/bhtoan2204/user/internal/application/command/command"
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

func (h *GetUserByIdCommandHandler) Handle(cmd *command.GetUserByIdCommand) (*command.GetUserByIdCommandResult, error) {
	return h.userService.GetUserById(cmd)
}
