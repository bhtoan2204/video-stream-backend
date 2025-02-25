package query

import "github.com/bhtoan2204/user/internal/domain/entities"

type ESRefreshTokenRepository interface {
	Index(*entities.RefreshToken) error
}
