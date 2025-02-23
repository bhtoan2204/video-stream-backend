package query

import (
	"github.com/bhtoan2204/user/internal/application/query/query"
	"github.com/bhtoan2204/user/internal/domain/entities"
)

type ESUserRepository interface {
	Index(*entities.User) error
	Search(query *query.SearchUserQuery) (*query.SearchUserQueryResult, error)
}
