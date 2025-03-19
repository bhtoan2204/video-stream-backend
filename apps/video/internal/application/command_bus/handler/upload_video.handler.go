package handler

import (
	"context"

	"github.com/bhtoan2204/video/internal/application/command_bus/command"
	"github.com/bhtoan2204/video/internal/domain/interfaces"
)

type UploadVideoCommandHandler struct {
	videoService interfaces.VideoServiceInterface
}

func NewUploadVideoCommandHandler(videoService interfaces.VideoServiceInterface) *UploadVideoCommandHandler {
	return &UploadVideoCommandHandler{
		videoService: videoService,
	}
}

func (h *UploadVideoCommandHandler) Handle(ctx context.Context, cmd *command.UploadVideoCommand) (*command.UploadVideoCommandResult, error) {
	return h.videoService.UploadVideo(ctx, cmd)
}
