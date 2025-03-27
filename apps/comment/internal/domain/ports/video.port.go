package ports

import (
	"context"
)

type VideoPort interface {
	GetVideo(ctx context.Context, videoID string) (*Video, error)
}

type Video struct {
	Id           string
	Title        string
	Description  string
	IsSearchable bool
	IsPublic     bool
	VideoUrl     string
}
