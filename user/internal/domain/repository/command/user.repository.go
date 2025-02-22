package command

import (
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/utils"
)

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindOneByQuery(query utils.QueryOptions) (*entities.User, error)
	FindByQuery(query utils.QueryOptions) ([]entities.User, error)
}
