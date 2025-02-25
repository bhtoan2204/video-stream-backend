package model

import "time"

type ESRefreshToken struct {
	ESAbstractModel
	Token     string     `json:"token"`
	UserID    string     `json:"user_id"`
	ExpiresAt time.Time  `json:"expires_at"`
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}
