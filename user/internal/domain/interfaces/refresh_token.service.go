package interfaces

import "time"

type RefreshTokenServiceInterface interface {
	CreateRefreshToken(token string, userId string, expires_at time.Time) error
	HardDeleteByQuery(query map[string]interface{}) error
	RevokedByQuery(query map[string]interface{}) error
}
