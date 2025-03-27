package handler

import (
	"context"

	"github.com/bhtoan2204/video/internal/application/command_bus/command"
	"github.com/bhtoan2204/video/internal/domain/interfaces"
)

type GetVideoByURLCommandHandler struct {
	videoService interfaces.VideoServiceInterface
}

func NewGetVideoByURLCommandHandler(videoService interfaces.VideoServiceInterface) *GetVideoByURLCommandHandler {
	return &GetVideoByURLCommandHandler{
		videoService: videoService,
	}
}

func (h *GetVideoByURLCommandHandler) Handle(ctx context.Context, cmd *command.GetVideoByURLCommand) (*command.GetVideoByURLCommandResult, error) {
	return h.videoService.GetVideoByURL(ctx, cmd)
}
