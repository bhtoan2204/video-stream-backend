package model_test

type ActivityLog struct {
	AbstractModel
	UserID      string `json:"user_id" gorm:"index;not null"`
	Action      string `json:"action" gorm:"type:varchar(255);not null"`
	Description string `json:"description,omitempty" gorm:"type:varchar(255)"`
	IPAddress   string `json:"ip_address,omitempty" gorm:"type:varchar(255)"`
	UserAgent   string `json:"user_agent,omitempty" gorm:"type:varchar(255)"`
	DeviceToken string `json:"device_token" gorm:"type:varchar(255)"`
}
