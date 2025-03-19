package interfaces

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/command_bus/command"
)

type UserSettingServiceInterface interface {
	UpdateByUserId(ctx context.Context, userId string, cmd *command.UpdateUserSettingsCommand) (*command.UpdateUserSettingsCommandResult, error)
}
