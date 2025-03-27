package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// LogoutRequest represents logout request
// @Description Logout request payload
type LogoutRequest struct {
	// Refresh token
	// @Example eyJhbGciOiJIUzI1...
	RefreshToken string `json:"refresh_token" validate:"required" example:"eyJhbGciOiJIUzI1..."`
}

func (c *LogoutRequest) Validate() error {
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
