package query

import (
	"context"

	"github.com/bhtoan2204/user/internal/domain/entities"
)

type ESRefreshTokenRepositoryInterface interface {
	Index(ctx context.Context, refresh_token *entities.RefreshToken) error
}
