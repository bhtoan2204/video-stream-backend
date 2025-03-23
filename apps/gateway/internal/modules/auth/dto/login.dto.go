package dto

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	TOTP     string `json:"totp"`
}

func (c *LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
