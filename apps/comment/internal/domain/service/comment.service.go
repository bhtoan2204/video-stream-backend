package service

import (
	"context"

	"github.com/bhtoan2204/comment/internal/application/command_bus/command"
	"github.com/bhtoan2204/comment/internal/domain/entities"
	repository_interface "github.com/bhtoan2204/comment/internal/domain/repository/command"
)

type CommentService struct {
	commentRepository repository_interface.CommentRepositoryInterface
}

func NewCommentService(commentRepository repository_interface.CommentRepositoryInterface) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
	}
}

func (s *CommentService) CreateComment(ctx context.Context, cmd *command.CreateCommentCommand) (*command.CreateCommentCommandResult, error) {
	var comment *entities.Comment

	comment = entities.NewComment(cmd.VideoId, cmd.UserId, cmd.ParentID, cmd.Content)

	result, err := s.commentRepository.Create(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &command.CreateCommentCommandResult{
		Result: result,
	}, nil
}
