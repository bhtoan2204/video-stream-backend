package query

import "github.com/bhtoan2204/user/internal/domain/entities"

type ESPermissionRepository interface {
	Index(*entities.Permission) error
}
