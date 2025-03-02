package command

import "github.com/go-playground/validator/v10"

type LogoutCommand struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (*LogoutCommand) CommandName() string {
	return "LogoutCommand"
}

func (c *LogoutCommand) Validate() error {
	validator := validator.New()
	return validator.Struct(c)
}
