package command

import "github.com/go-playground/validator/v10"

type UpdateUserSettingsCommand struct {
	UserID               string                 `json:"user_id" validate:"required"`
	Language             *string                `json:"language,omitempty"`
	Theme                *string                `json:"theme,omitempty"`
	NotificationsEnabled *bool                  `json:"notifications_enabled,omitempty"`
	Is2FAEnabled         *bool                  `json:"is_2fa_enabled,omitempty"`
	Privacy              *PrivacySettingsUpdate `json:"privacy,omitempty"`
}

type UpdateUserSettingsCommandResult struct {
	Message string `json:"message"`
}

type PrivacySettingsUpdate struct {
	ShowEmail       *bool `json:"show_email,omitempty"`
	ShowSubscribers *bool `json:"show_subscribers,omitempty"`
}

func (c *UpdateUserSettingsCommand) CommandName() string {
	return "UpdateUserSettingsCommand"
}

func (c *UpdateUserSettingsCommand) Validate() error {
	validator := validator.New()
	return validator.Struct(c)
}
