package command

import (
	common "github.com/bhtoan2204/user/internal/application/common/command"
	"github.com/go-playground/validator/v10"
)

type LoginCommand struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	TOTP     string `json:"totp"`
}

type LoginCommandResult struct {
	Result *common.LoginResult `json:"result"`
}

func (*LoginCommand) CommandName() string {
	return "LoginCommand"
}

func (c *LoginCommand) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
