package entities

import (
	"time"
)

type RefreshToken struct {
	AbstractModel
	Token     string     `json:"token" gorm:"uniqueIndex;not null"`
	UserID    string     `json:"user_id" gorm:"index;not null"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null"`
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

type IRefreshToken interface {
}
