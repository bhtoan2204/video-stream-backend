package persistent_object

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
	Theme                Theme            `json:"theme,omitempty" gorm:"type:enum('light','dark');default:'light'"`
	NotificationsEnabled bool             `json:"notifications_enabled,omitempty"`
	Privacy              *PrivacySettings `json:"privacy,omitempty" gorm:"embedded"`
}

func (UserSettings) TableName() string {
	return "user_settings"
}
