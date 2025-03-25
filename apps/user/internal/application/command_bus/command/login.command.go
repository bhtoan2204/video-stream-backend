package command

import (
	"fmt"

	common "github.com/bhtoan2204/user/internal/application/common/command"
	"github.com/go-playground/validator/v10"
)

type LoginCommand struct {
	Email    string `json:"email" validate:"required"`
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
