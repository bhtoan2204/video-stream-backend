package query

import "github.com/bhtoan2204/user/internal/domain/entities"

type ESActivityLogRepository interface {
	Index(*entities.ActivityLog) error
}
