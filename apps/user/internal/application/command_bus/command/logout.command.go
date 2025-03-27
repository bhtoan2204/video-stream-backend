package command

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type LogoutCommand struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (*LogoutCommand) CommandName() string {
	return "LogoutCommand"
}

func (c *LogoutCommand) Validate() error {
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
