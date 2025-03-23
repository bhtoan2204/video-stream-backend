package dto

import "github.com/go-playground/validator/v10"

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (c *LogoutRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
