package command

import common "github.com/bhtoan2204/user/internal/application/common/command"

type CreateUserCommand struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required"`
}

type CreateUserCommandResult struct {
	Result *common.UserResult `json:"result"`
}

func (*CreateUserCommand) CommandName() string {
	return "CreateUserCommand"
}
