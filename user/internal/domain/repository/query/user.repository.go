package query

import "github.com/bhtoan2204/user/internal/domain/entities"

type ESUserRepository interface {
	Index(*entities.User) error
	Search(query string) ([]*entities.User, error)
}
