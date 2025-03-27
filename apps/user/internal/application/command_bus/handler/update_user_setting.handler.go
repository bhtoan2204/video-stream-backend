package handler

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/command_bus/command"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
)

type UpdateUserSettingCommandHandler struct {
	userSettingService interfaces.UserSettingServiceInterface
}

func NewUpdateUserSettingCommandHandler(userSettingService interfaces.UserSettingServiceInterface) *UpdateUserSettingCommandHandler {
	return &UpdateUserSettingCommandHandler{
		userSettingService: userSettingService,
	}
}

func (h *UpdateUserSettingCommandHandler) Handle(ctx context.Context, cmd *command.UpdateUserSettingsCommand) (*command.UpdateUserSettingsCommandResult, error) {
	return h.userSettingService.UpdateByUserId(ctx, cmd.UserID, cmd)
}
