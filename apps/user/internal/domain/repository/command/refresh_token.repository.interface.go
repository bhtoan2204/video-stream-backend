package repository_interface

import (
	"context"

	"github.com/bhtoan2204/user/internal/domain/entities"
)

type RefreshTokenRepositoryInterface interface {
	Create(ctx context.Context, refreshToken *entities.RefreshToken) error
	FindOneByQuery(ctx context.Context, query *map[string]interface{}) (*entities.RefreshToken, error)
	DeleteByQuery(ctx context.Context, query *map[string]interface{}) error
	HardDeleteByQuery(ctx context.Context, query *map[string]interface{}) error
	UpdateByQuery(ctx context.Context, query *map[string]interface{}, update *map[string]interface{}) error
}
