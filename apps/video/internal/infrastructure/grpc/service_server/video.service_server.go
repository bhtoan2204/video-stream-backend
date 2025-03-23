package service_server

import (
	"context"

	repository "github.com/bhtoan2204/video/internal/domain/repository/command"
	gprcVideo "github.com/bhtoan2204/video/internal/infrastructure/grpc/proto/video"
	"github.com/bhtoan2204/video/utils"
)

type VideoServiceServerImpl struct {
	gprcVideo.UnimplementedVideoServiceServer
	videoRepository repository.VideoRepositoryInterface
}

func NewVideoServiceServer(videoRepository repository.VideoRepositoryInterface) gprcVideo.VideoServiceServer {
	return &VideoServiceServerImpl{
		videoRepository: videoRepository,
	}
}

func (s *VideoServiceServerImpl) GetVideo(ctx context.Context, req *gprcVideo.GetVideoRequest) (*gprcVideo.GetVideoResponse, error) {
	video, err := s.videoRepository.FindOne(ctx, &utils.QueryOptions{
		Filters: map[string]interface{}{"id": req.Id},
	})

	if err != nil {
		return nil, err
	}

	return &gprcVideo.GetVideoResponse{
		Id:           video.ID,
		Title:        video.Title,
		Description:  video.Description,
		IsSearchable: video.IsSearchable,
		IsPublic:     video.IsPublic,
		VideoUrl:     video.VideoURL,
	}, nil
}

func (s *VideoServiceServerImpl) mustEmbedUnimplementedVideoServiceServer() {}
