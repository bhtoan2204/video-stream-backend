package query

import (
	"context"

	"github.com/bhtoan2204/user/internal/domain/entities"
)

type ESPermissionRepositoryInterface interface {
	Index(ctx context.Context, permission *entities.Permission) error
}
