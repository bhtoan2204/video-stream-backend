package command

import (
	"context"

	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/utils"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	FindOneByQuery(ctx context.Context, query *utils.QueryOptions) (*entities.User, error)
	FindByQuery(ctx context.Context, query *utils.QueryOptions) ([]entities.User, error)
}
