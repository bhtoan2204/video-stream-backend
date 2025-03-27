package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// LoginRequest represents login credentials
// @Description Login request payload
type LoginRequest struct {
	// User email
	// @Example john@example.com
	Email string `json:"email" validate:"required" example:"john@example.com"`
	// User password
	// @Example secretpass123
	Password string `json:"password" validate:"required,min=8" example:"secretpass123"`
	// Two-factor authentication code (optional)
	// @Example 123456
	TOTP string `json:"totp" example:"123456"`
}

func (c *LoginRequest) Validate() error {
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
