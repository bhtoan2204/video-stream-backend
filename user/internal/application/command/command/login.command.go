package command

import common "github.com/bhtoan2204/user/internal/application/common/command"

type LoginCommand struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginCommandResult struct {
	Result *common.LoginResult `json:"result"`
}

func (*LoginCommand) CommandName() string {
	return "LoginCommand"
}
