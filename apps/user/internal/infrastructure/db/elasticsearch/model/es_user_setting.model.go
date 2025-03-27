package model

type ESPrivacySetting struct {
	ShowEmail       bool `json:"show_email"`
	ShowSubscribers bool `json:"show_subscribers"`
}

type ESUserSetting struct {
	ESAbstractModel
	UserID               string            `json:"user_id"`
	Language             string            `json:"language,omitempty"`
	Theme                string            `json:"theme,omitempty"` // "light" or "dark"
	NotificationsEnabled bool              `json:"notifications_enabled"`
	Privacy              *ESPrivacySetting `json:"privacy,omitempty"`
}
