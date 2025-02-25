package entities

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
	UserID               string           `json:"user_id"`
	Language             string           `json:"language,omitempty"`
	Theme                Theme            `json:"theme,omitempty"`
	NotificationsEnabled bool             `json:"notifications_enabled,omitempty"`
	Privacy              *PrivacySettings `json:"privacy,omitempty"`
}

type IUserSettings interface {
}
