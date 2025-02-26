package manager

import "time"

type RefreshTokenManager interface {
	CreateRefreshToken(token string, userId string, expiresAt time.Time) error
	HardDeleteByQuery(query map[string]interface{}) error
	RevokedByQuery(query map[string]interface{}) error
	CheckRefreshToken(token string) (bool, error)
}
