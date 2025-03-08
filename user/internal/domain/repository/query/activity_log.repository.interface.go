package es_repository_interface

import (
	"context"

	"github.com/bhtoan2204/user/internal/domain/entities"
)

type ESActivityLogRepositoryInterface interface {
	Index(ctx context.Context, activity_log *entities.ActivityLog) error
}
