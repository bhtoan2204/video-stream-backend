package query

import (
	"context"

	"github.com/bhtoan2204/user/internal/domain/entities"
)

type ESUserSettingRepositoryInterface interface {
	Index(ctx context.Context, user_settings *entities.UserSettings) error
}
