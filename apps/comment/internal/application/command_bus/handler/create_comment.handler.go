package handler

import (
	"context"

	"github.com/bhtoan2204/comment/internal/application/command_bus/command"
	"github.com/bhtoan2204/comment/internal/domain/interfaces"
)

type CreateCommentCommandHandler struct {
	commentService interfaces.CommentServiceInterface
}

func NewCreateCommentCommandHandler(commentService interfaces.CommentServiceInterface) *CreateCommentCommandHandler {
	return &CreateCommentCommandHandler{
		commentService: commentService,
	}
}

func (h *CreateCommentCommandHandler) Handle(ctx context.Context, cmd *command.CreateCommentCommand) (*command.CreateCommentCommandResult, error) {
	return h.commentService.CreateComment(ctx, cmd)
}
