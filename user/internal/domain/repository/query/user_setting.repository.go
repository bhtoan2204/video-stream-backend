package query

import "github.com/bhtoan2204/user/internal/domain/entities"

type ESUserSettingRepository interface {
	Index(*entities.UserSettings) error
}
