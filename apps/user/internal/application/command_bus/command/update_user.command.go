package command

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UpdateUserCommand struct {
	ID        string `json:"id" validate:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	BirthDate string `json:"birth_date"`
	Avatar    string `json:"avatar"`
}

type UpdateUserCommandResult struct {
	Success bool `json:"success"`
}

func (c *UpdateUserCommand) Validate() error {
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

func (c *UpdateUserCommand) CommandName() string {
	return "UpdateUserCommand"
}
