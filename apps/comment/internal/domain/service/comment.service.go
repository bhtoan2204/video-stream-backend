package service

import (
	"context"
	"fmt"

	"github.com/bhtoan2204/comment/internal/application/command_bus/command"
	"github.com/bhtoan2204/comment/internal/domain/entities"
	"github.com/bhtoan2204/comment/internal/domain/ports"
	repository_interface "github.com/bhtoan2204/comment/internal/domain/repository/command"
)

type CommentService struct {
	commentRepository repository_interface.CommentRepositoryInterface
	videoPort         ports.VideoPort
}

func NewCommentService(
	commentRepository repository_interface.CommentRepositoryInterface,
	videoPort ports.VideoPort,
) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
		videoPort:         videoPort,
	}
}

func (s *CommentService) CreateComment(ctx context.Context, cmd *command.CreateCommentCommand) (*command.CreateCommentCommandResult, error) {
	_, err := s.videoPort.GetVideo(ctx, cmd.VideoId)
	if err != nil {
		return nil, fmt.Errorf("failed to validate video: %w", err)
	}

	err = cmd.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate comment: %w", err)
	}
	var comment *entities.Comment

	comment = entities.NewComment(cmd.VideoId, cmd.UserId, cmd.ParentID, cmd.Content)
	fmt.Println("comment......", comment)
	result, err := s.commentRepository.Create(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &command.CreateCommentCommandResult{
		Result: result,
	}, nil
}
