package entities

import (
	"time"
)

type RefreshToken struct {
	AbstractModel
	Token     string     `json:"token"`
	UserID    string     `json:"user_id"`
	ExpiresAt time.Time  `json:"expires_at"`
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}
