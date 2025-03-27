package es_repository_interface

import (
	"context"

	"github.com/bhtoan2204/user/internal/domain/entities"
)

type ESRoleRepositoryInterface interface {
	Index(ctx context.Context, role *entities.Role) error
}
