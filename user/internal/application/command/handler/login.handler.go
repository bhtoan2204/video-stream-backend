package handler

import (
	"github.com/bhtoan2204/user/internal/application/command/command"
	"github.com/bhtoan2204/user/internal/application/interfaces"
)

type LoginCommandHandler struct {
	userService interfaces.UserServiceInterface
}

func NewLoginCommandHandler(userService interfaces.UserServiceInterface) *LoginCommandHandler {
	return &LoginCommandHandler{
		userService: userService,
	}
}

func (h *LoginCommandHandler) Handle(cmd *command.LoginCommand) (*command.LoginCommandResult, error) {
	return h.userService.Login(cmd)
}
