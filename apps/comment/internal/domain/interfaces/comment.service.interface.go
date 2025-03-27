package interfaces

import (
	"context"

	"github.com/bhtoan2204/comment/internal/application/command_bus/command"
)

type CommentServiceInterface interface {
	CreateComment(ctx context.Context, cmd *command.CreateCommentCommand) (*command.CreateCommentCommandResult, error)
}
