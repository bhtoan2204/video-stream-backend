package handler

import (
	"context"

	"github.com/bhtoan2204/video/internal/application/command_bus/command"
	"github.com/bhtoan2204/video/internal/domain/interfaces"
)

type DeleteVideoCommandHandler struct {
	videoService interfaces.VideoServiceInterface
}

func NewDeleteVideoCommandHandler(videoService interfaces.VideoServiceInterface) *DeleteVideoCommandHandler {
	return &DeleteVideoCommandHandler{
		videoService: videoService,
	}
}

func (h *DeleteVideoCommandHandler) Handle(ctx context.Context, cmd *command.DeleteVideoCommand) (*command.DeleteVideoCommandResult, error) {
	return h.videoService.DeleteVideo(ctx, cmd)
}
