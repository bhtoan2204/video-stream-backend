package query

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/query/query"
	"github.com/bhtoan2204/user/internal/domain/entities"
)

type ESUserRepositoryInterface interface {
	Index(ctx context.Context, user *entities.User) error
	Search(ctx context.Context, query *query.SearchUserQuery) (*[]entities.User, *query.PaginateResult, error)
}
