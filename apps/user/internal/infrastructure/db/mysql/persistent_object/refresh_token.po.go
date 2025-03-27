package persistent_object

import (
	"time"
)

type RefreshToken struct {
	BasePO
	Token     string     `json:"token" gorm:"type:varchar(512);uniqueIndex;not null"`
	UserID    string     `json:"user_id" gorm:"type:varchar(255);index;not null"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null"`
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
