package command

import (
	common "github.com/bhtoan2204/user/internal/application/common/command"
	"github.com/go-playground/validator/v10"
)

type CreateUserCommand struct {
	Username  string `json:"username" validate:"required,min=3,max=20"`
	Password  string `json:"password" validate:"required,min=8"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
	Address   string `json:"address" validate:"required"`
}

type CreateUserCommandResult struct {
	Result *common.UserResult `json:"result"`
}

func (*CreateUserCommand) CommandName() string {
	return "CreateUserCommand"
}

func (c *CreateUserCommand) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
