package query

import "github.com/bhtoan2204/user/internal/domain/entities"

type ESRoleRepository interface {
	Index(*entities.Role) error
}
