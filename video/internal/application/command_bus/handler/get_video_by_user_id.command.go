package handler

import (
	"context"

	"github.com/bhtoan2204/video/internal/application/command_bus/command"
	"github.com/bhtoan2204/video/internal/domain/interfaces"
)

type GetVideoByUserIdCommandHandler struct {
	videoService interfaces.VideoServiceInterface
}

func NewGetVideoByUserIdCommandHandler(videoService interfaces.VideoServiceInterface) *GetVideoByUserIdCommandHandler {
	return &GetVideoByUserIdCommandHandler{
		videoService: videoService,
	}
}

func (h *GetVideoByUserIdCommandHandler) Handle(ctx context.Context, cmd *command.GetVideoByUserIdCommand) (*command.GetVideoByUserIdCommandResult, error) {
	return h.videoService.GetVideoByUserId(ctx, cmd)
}
