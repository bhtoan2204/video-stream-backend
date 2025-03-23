package command_bus

import (
	"context"

	"github.com/bhtoan2204/comment/internal/application/command_bus/command"
	"github.com/bhtoan2204/comment/internal/application/command_bus/handler"
	"github.com/bhtoan2204/comment/internal/application/shared"
)

func SetUpCommandBus(deps *shared.ServiceDependencies) *CommandBus {
	bus := NewCommandBus()

	// Register User command handlers
	createCommentHandler := handler.NewCreateCommentCommandHandler(deps.CommentService)

	bus.RegisterHandler("CreateCommentCommand", func(ctx context.Context, cmd Command) (interface{}, error) {
		return createCommentHandler.Handle(ctx, cmd.(*command.CreateCommentCommand))
	})

	return bus
}
