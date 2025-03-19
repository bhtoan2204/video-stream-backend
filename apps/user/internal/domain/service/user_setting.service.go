package service

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/command_bus/command"
	repository_interface "github.com/bhtoan2204/user/internal/domain/repository/command"
)

type UserSettingService struct {
	userSettingRepository repository_interface.UserSettingRepositoryInterface
}

func NewUserSettingService(userSettingRepository repository_interface.UserSettingRepositoryInterface) *UserSettingService {
	return &UserSettingService{
		userSettingRepository: userSettingRepository,
	}
}

func (s *UserSettingService) UpdateByUserId(ctx context.Context, userId string, cmd *command.UpdateUserSettingsCommand) (*command.UpdateUserSettingsCommandResult, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	update := map[string]interface{}{}

	if cmd.Language != nil {
		update["language"] = *cmd.Language
	}

	if cmd.Theme != nil {
		update["theme"] = *cmd.Theme
	}

	if cmd.NotificationsEnabled != nil {
		update["notifications_enabled"] = *cmd.NotificationsEnabled
	}

	if cmd.Is2FAEnabled != nil {
		update["is_2fa_enabled"] = *cmd.Is2FAEnabled
	}

	err := s.userSettingRepository.UpdateByUserId(ctx, &userId, &update)
	if err != nil {
		return nil, err
	}

	return &command.UpdateUserSettingsCommandResult{
		Message: "Update user setting successfully",
	}, nil
}
