package command

import "github.com/go-playground/validator/v10"

type RefreshTokenCommand struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (*RefreshTokenCommand) CommandName() string {
	return "RefreshTokenCommand"
}

func (c *RefreshTokenCommand) Validate() error {
	validator := validator.New()
	return validator.Struct(c)
}
