package repository

import (
	"github.com/bhtoan2204/user/internal/application/query"
	"github.com/bhtoan2204/user/internal/domain/entities"
)

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindOneByQuery(query query.QueryOptions) (*entities.User, error)
	FindByQuery(query query.QueryOptions) ([]entities.User, error)
}
