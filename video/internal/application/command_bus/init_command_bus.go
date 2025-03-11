package command_bus

import (
	"context"

	"github.com/bhtoan2204/video/internal/application/command_bus/command"
	"github.com/bhtoan2204/video/internal/application/command_bus/handler"
	"github.com/bhtoan2204/video/internal/application/shared"
)

func SetUpCommandBus(deps *shared.ServiceDependencies) *CommandBus {
	bus := NewCommandBus()

	deleteVideoHandler := handler.NewDeleteVideoCommandHandler(deps.VideoService)
	createVideoHandler := handler.NewUploadVideoCommandHandler(deps.VideoService)
	getVideoByIdHandler := handler.NewGetVideoByIdCommandHandler(deps.VideoService)
	getVideoByUserIdHandler := handler.NewGetVideoByUserIdCommandHandler(deps.VideoService)

	bus.RegisterHandler("DeleteVideoCommand", func(ctx context.Context, cmd Command) (interface{}, error) {
		return deleteVideoHandler.Handle(ctx, cmd.(*command.DeleteVideoCommand))
	})
	bus.RegisterHandler("UploadVideoCommand", func(ctx context.Context, cmd Command) (interface{}, error) {
		return createVideoHandler.Handle(ctx, cmd.(*command.UploadVideoCommand))
	})
	bus.RegisterHandler("GetVideoByIdCommand", func(ctx context.Context, cmd Command) (interface{}, error) {
		return getVideoByIdHandler.Handle(ctx, cmd.(*command.GetVideoByIdCommand))
	})
	bus.RegisterHandler("GetVideoByUserIdCommand", func(ctx context.Context, cmd Command) (interface{}, error) {
		return getVideoByUserIdHandler.Handle(ctx, cmd.(*command.GetVideoByUserIdCommand))
	})

	return bus
}
