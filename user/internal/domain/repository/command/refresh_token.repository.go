package command

import "github.com/bhtoan2204/user/internal/domain/entities"

type RefreshTokenRepository interface {
	Create(refreshToken *entities.RefreshToken) error
	FindOneByQuery(query map[string]interface{}) (*entities.RefreshToken, error)
	DeleteByQuery(query map[string]interface{}) error
	UpdateByQuery(query map[string]interface{}, update map[string]interface{}) error
}
