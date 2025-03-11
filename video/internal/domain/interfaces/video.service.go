package interfaces

import (
	"context"

	"github.com/bhtoan2204/video/internal/application/command_bus/command"
)

type VideoServiceInterface interface {
	UploadVideo(ctx context.Context, uploadVideoCommand *command.UploadVideoCommand) (*command.UploadVideoCommandResult, error)
	GetVideoById(ctx context.Context, getVideoByIdCommand *command.GetVideoByIdCommand) (*command.GetVideoByIdCommandResult, error)
	GetVideoByUserId(ctx context.Context, getVideoByUserIdCommand *command.GetVideoByUserIdCommand) (*command.GetVideoByUserIdCommandResult, error)
	DeleteVideo(ctx context.Context, deleteVideoCommand *command.DeleteVideoCommand) (*command.DeleteVideoCommandResult, error)
}
