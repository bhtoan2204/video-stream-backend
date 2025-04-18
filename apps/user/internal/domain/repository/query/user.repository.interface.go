package es_repository_interface

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/query_bus/query"
	"github.com/bhtoan2204/user/internal/domain/entities"
)

type ESUserRepositoryInterface interface {
	Index(ctx context.Context, user *entities.User) error
	Search(ctx context.Context, query *query.SearchUserQuery) (*[]entities.User, *query.PaginateResult, error)
}
