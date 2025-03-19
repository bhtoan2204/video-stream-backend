package persistent_object_test

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
	BasePO
	UserID               string           `json:"user_id" gorm:"index;not null"`
	Language             string           `json:"language,omitempty"`
	Theme                string           `json:"theme" gorm:"type:text;default:'light'"`
	NotificationsEnabled bool             `json:"notifications_enabled,omitempty"`
	Privacy              *PrivacySettings `json:"privacy,omitempty" gorm:"embedded"`
	Is2FAEnabled         bool             `json:"is_2fa_enabled" gorm:"default:false"`
	TOTPSecret           string           `json:"-" gorm:"column:totp_secret"`
}

func (UserSettings) TableName() string {
	return "user_settings"
}
