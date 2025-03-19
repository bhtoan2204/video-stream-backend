package model

import (
	"time"

	"gorm.io/gorm"
)

type Reaction struct {
	CommentID string         `json:"comment_id" gorm:"type:char(36);not null;primaryKey"`
	UserID    string         `json:"user_id" gorm:"type:char(36);not null;primaryKey"`
	IsLike    bool           `json:"like,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
