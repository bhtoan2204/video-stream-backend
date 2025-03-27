package model

import (
	"time"

	"github.com/google/uuid"
)

type Video struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description,omitempty"`
	IsSearchable bool      `json:"is_searchable,omitempty"`
	IsPublic     bool      `json:"is_public,omitempty"`
	VideoURL     string    `json:"video_url"`
	Bucket       string    `json:"bucket"`
	ObjectKey    string    `json:"object_key"`
	UploadedUser string    `json:"uploaded_user"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
