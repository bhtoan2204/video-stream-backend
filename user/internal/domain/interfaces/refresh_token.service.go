package interfaces

import (
	"context"
	"time"
)

type RefreshTokenServiceInterface interface {
	CreateRefreshToken(ctx context.Context, token string, userId string, expires_at time.Time) error
	HardDeleteByQuery(ctx context.Context, query map[string]interface{}) error
	RevokedByQuery(ctx context.Context, query map[string]interface{}) error
	CheckRefreshToken(ctx context.Context, token string) (bool, error)
}
