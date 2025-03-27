package command

import (
	"fmt"

	common "github.com/bhtoan2204/user/internal/application/common/command"
	"github.com/go-playground/validator/v10"
)

type CreateUserCommand struct {
	Username  string `json:"username" validate:"required,min=3,max=20"`
	Password  string `json:"password" validate:"required,min=8"`
	Email     string `json:"email" validate:"required"`
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
	err := validate.Struct(c)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		var errorMessage string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage += fmt.Sprintf("Field: %s, Error: %s, Value: %v\n",
				err.Field(),
				err.Tag(),
				err.Value())
		}
		return fmt.Errorf(errorMessage)
	}
	return nil
}
