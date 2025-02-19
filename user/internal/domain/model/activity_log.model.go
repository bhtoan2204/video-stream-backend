package model

import "gorm.io/gorm"

type ActivityLog struct {
	AbstractModel
	UserID      string         `json:"user_id" gorm:"index"`
	Action      string         `json:"action"`
	Description string         `json:"description,omitempty"`
	IPAddress   string         `json:"ip_address,omitempty"`
	UserAgent   string         `json:"user_agent,omitempty"`
	DeviceToken string         `json:"device_token"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
