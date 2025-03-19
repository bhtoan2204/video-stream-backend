package command

import "github.com/go-playground/validator/v10"

type Setup2FACommand struct {
	UserID string `json:"user_id" validate:"required"`
}

type Setup2FACommandResult struct {
	OTPUrl string `json:"otp_url"`
	QRCode string `json:"qr_code"`
}

func (*Setup2FACommand) CommandName() string {
	return "Setup2FACommand"
}

func (c *Setup2FACommand) Validate() error {
	validator := validator.New()
	return validator.Struct(c)
}
