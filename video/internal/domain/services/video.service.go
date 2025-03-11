package services

import (
	"context"

	"github.com/bhtoan2204/video/internal/application/command_bus/command"
	repository_interface "github.com/bhtoan2204/video/internal/domain/repository/command"
)

type VideoService struct {
	videoRepository repository_interface.VideoRepositoryInterface
}

func NewVideoService(videoRepository repository_interface.VideoRepositoryInterface) *VideoService {
	return &VideoService{
		videoRepository: videoRepository,
	}
}

func (r *VideoService) UploadVideo(ctx context.Context, uploadVideoCommand *command.UploadVideoCommand) (*command.UploadVideoCommandResult, error) {
	panic("not implemented") // TODO: Implement
}

func (r *VideoService) GetVideoById(ctx context.Context, getVideoByIdCommand *command.GetVideoByIdCommand) (*command.GetVideoByIdCommandResult, error) {
	panic("not implemented") // TODO: Implement
}

func (r *VideoService) GetVideoByUserId(ctx context.Context, getVideoByUserIdCommand *command.GetVideoByUserIdCommand) (*command.GetVideoByUserIdCommandResult, error) {
	panic("not implemented") // TODO: Implement
}

func (r *VideoService) DeleteVideo(ctx context.Context, deleteVideoCommand *command.DeleteVideoCommand) (*command.DeleteVideoCommandResult, error) {
	panic("not implemented") // TODO: Implement
}
