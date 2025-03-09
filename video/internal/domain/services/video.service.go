package services

import (
	"context"

	"github.com/bhtoan2204/video/internal/application/command/command"
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

func UploadVideo(ctx context.Context, uploadVideoCommand *command.UploadVideoCommand) (*command.UploadVideoCommandResult, error) {
	return nil, nil
}
func GetVideoById(ctx context.Context, getVideoByIdCommand *command.GetVideoByIdCommand) (*command.GetVideoByIdCommandResult, error) {
	return nil, nil
}
func GetVideoByUserId(ctx context.Context, getVideoByUserIdCommand *command.GetVideoByUserIdCommand) (*command.GetVideoByUserIdCommandResult, error) {
	return nil, nil
}
func DeleteVideo(ctx context.Context, deleteVideoCommand *command.DeleteVideoCommand) (*command.DeleteVideoCommandResult, error) {
	return nil, nil
}
