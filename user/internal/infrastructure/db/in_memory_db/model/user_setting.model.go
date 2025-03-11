package model_test

type Theme string

const (
	ThemeLight Theme = "light"
	ThemeDark  Theme = "dark"
)

type PrivacySettings struct {
	ShowEmail       bool `json:"show_email,omitempty"`
	ShowSubscribers bool `json:"show_subscribers,omitempty"`
}

type UserSettings struct {
	AbstractModel
	UserID               string           `json:"user_id" gorm:"index;not null"`
	Language             string           `json:"language,omitempty"`
	Theme                Theme            `json:"theme,omitempty" gorm:"type:enum('light','dark');default:'light'"`
	NotificationsEnabled bool             `json:"notifications_enabled,omitempty"`
	Privacy              *PrivacySettings `json:"privacy,omitempty" gorm:"embedded"`
	Is2FAEnabled         bool             `json:"is_2fa_enabled,omitempty"`
	TOTPSecret           string           `json:"-" gorm:"column:totp_secret"`
}
