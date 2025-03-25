package service

import (
	"context"
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/bhtoan2204/comment/global"
	"github.com/bhtoan2204/comment/internal/application/command_bus/command"
	"github.com/bhtoan2204/comment/internal/domain/entities"
	"github.com/bhtoan2204/comment/internal/domain/ports"
	repository_interface "github.com/bhtoan2204/comment/internal/domain/repository/command"
	"go.uber.org/zap"
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
	var video *ports.Video
	err := hystrix.Do("get_video_command", func() error {
		var err error
		video, err = s.videoPort.GetVideo(ctx, cmd.VideoId)
		if err != nil {
			return fmt.Errorf("failed to get video: %w", err)
		}
		return nil
	}, func(fallbackErr error) error {
		// Log the fallback event and return a custom fallback error.
		global.Logger.Error("Hystrix fallback triggered for get_video_command", zap.Error(fallbackErr))
		return fmt.Errorf("fallback: unable to retrieve video information")
	})
	if err != nil {
		return nil, err
	}
	if video == nil {
		return nil, fmt.Errorf("video not found")
	}

	if err := cmd.Validate(); err != nil {
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
