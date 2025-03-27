package services

import (
	"context"
	"fmt"

	"github.com/bhtoan2204/video/global"
	"github.com/bhtoan2204/video/internal/application/command_bus/command"
	"github.com/bhtoan2204/video/internal/domain/entities"
	repository_interface "github.com/bhtoan2204/video/internal/domain/repository/command"
	"github.com/bhtoan2204/video/internal/infrastructure/grpc/proto/user"
	"github.com/bhtoan2204/video/utils"
	"github.com/hashicorp/go-uuid"
	"go.uber.org/zap"
)

type VideoService struct {
	videoRepository repository_interface.VideoRepositoryInterface
}

func NewVideoService(videoRepository repository_interface.VideoRepositoryInterface) *VideoService {
	return &VideoService{
		videoRepository: videoRepository,
	}
}

func (r *VideoService) UploadVideo(ctx context.Context, cmd *command.UploadVideoCommand) (*command.UploadVideoCommandResult, error) {
	if err := global.S3Client.VerifyFileExists(ctx, cmd.FileKey); err != nil {
		return nil, err
	}
	if err := cmd.Validate(); err != nil {
		return nil, err
	}
	videoUrl, _ := uuid.GenerateUUID()

	videoEntities := &entities.Video{
		Title:        cmd.Title,
		Description:  cmd.Description,
		IsSearchable: cmd.IsSearchable,
		IsPublic:     cmd.IsPublic,
		VideoURL:     videoUrl,
		Bucket:       global.Config.S3Config.Bucket,
		ObjectKey:    cmd.FileKey,
		UploadedUser: cmd.UploadedUser,
	}
	video, err := r.videoRepository.CreateOne(ctx, videoEntities)
	if err != nil {
		global.Logger.Error("Failed to create video", zap.Error(err))
		return nil, err
	}
	return &command.UploadVideoCommandResult{
		Result: video,
	}, nil
}

func (r *VideoService) GetVideoByURL(ctx context.Context, cmd *command.GetVideoByURLCommand) (*command.GetVideoByURLCommandResult, error) {
	userVal := ctx.Value("user")
	if err := cmd.Validate(); err != nil {
		return nil, err
	}
	userResp, _ := userVal.(*user.UserResponse)

	fmt.Println(userResp)

	video, err := r.videoRepository.FindOne(ctx, &utils.QueryOptions{
		Filters: map[string]interface{}{
			"video_url": cmd.URL,
		},
	})

	if err != nil {
		return nil, err
	}
	return &command.GetVideoByURLCommandResult{
		Result: video,
	}, nil
}
func (r *VideoService) GetVideoByUserId(ctx context.Context, cmd *command.GetVideoByUserIdCommand) (*command.GetVideoByUserIdCommandResult, error) {
	panic("not implemented") // TODO: Implement
}

func (r *VideoService) DeleteVideo(ctx context.Context, cmd *command.DeleteVideoCommand) (*command.DeleteVideoCommandResult, error) {
	panic("not implemented") // TODO: Implement
}
