package handler

import (
	"context"

	"github.com/bhtoan2204/video/internal/application/command_bus/command"
	"github.com/bhtoan2204/video/internal/domain/interfaces"
)

type GetVideoByIdCommandHandler struct {
	videoService interfaces.VideoServiceInterface
}

func NewGetVideoByIdCommandHandler(videoService interfaces.VideoServiceInterface) *GetVideoByIdCommandHandler {
	return &GetVideoByIdCommandHandler{
		videoService: videoService,
	}
}

func (h *GetVideoByIdCommandHandler) Handle(ctx context.Context, cmd *command.GetVideoByIdCommand) (*command.GetVideoByIdCommandResult, error) {
	return h.videoService.GetVideoById(ctx, cmd)
}
