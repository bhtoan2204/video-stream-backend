package grpc

import (
	"context"
	"fmt"

	"github.com/bhtoan2204/comment/internal/domain/ports"
	"github.com/bhtoan2204/comment/internal/infrastructure/grpc/proto/video"
)

type videoAdapter struct {
	client video.VideoServiceClient
}

func NewVideoAdapter(client video.VideoServiceClient) ports.VideoPort {
	return &videoAdapter{
		client: client,
	}
}

func (a *videoAdapter) GetVideo(ctx context.Context, videoID string) (*ports.Video, error) {
	video, err := a.client.GetVideo(ctx, &video.GetVideoRequest{
		Id: videoID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get video: %w", err)
	}

	return &ports.Video{
		Id:           video.Id,
		Title:        video.Title,
		Description:  video.Description,
		IsSearchable: video.IsSearchable,
		IsPublic:     video.IsPublic,
		VideoUrl:     video.VideoUrl,
	}, nil
}
