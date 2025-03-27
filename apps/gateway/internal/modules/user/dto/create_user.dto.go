package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=20"`
	Password  string `json:"password" validate:"required,min=8"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
	Address   string `json:"address" validate:"required"`
}

func (c *CreateUserRequest) Validate() error {
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
