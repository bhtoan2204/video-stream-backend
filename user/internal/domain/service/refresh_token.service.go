package service

import (
	"time"

	"github.com/bhtoan2204/user/internal/domain/entities"
	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
)

type RefreshTokenService struct {
	refreshTokenRepository repository.RefreshTokenRepository
}

func NewRefreshTokenService(refreshTokenRepository repository.RefreshTokenRepository) *RefreshTokenService {
	return &RefreshTokenService{
		refreshTokenRepository: refreshTokenRepository,
	}
}

func (s *RefreshTokenService) CreateRefreshToken(token string, userId string, expires_at time.Time) error {
	return s.refreshTokenRepository.Create(&entities.RefreshToken{
		Token:     token,
		UserID:    userId,
		ExpiresAt: expires_at,
	})
}

func (s *RefreshTokenService) HardDeleteByQuery(query map[string]interface{}) error {
	return s.refreshTokenRepository.HardDeleteByQuery(query)
}

func (s *RefreshTokenService) RevokedByQuery(query map[string]interface{}) error {
	return s.refreshTokenRepository.UpdateByQuery(query, map[string]interface{}{
		"revoked_at": time.Now(),
	})
}

func (s *RefreshTokenService) DeleteByQuery(query map[string]interface{}) error {
	return s.refreshTokenRepository.DeleteByQuery(query)
}

func (s *RefreshTokenService) CheckRefreshToken(token string) (bool, error) {
	refreshToken, err := s.refreshTokenRepository.FindOneByQuery(map[string]interface{}{
		"token": token,
		"revoked_at": map[string]interface{}{
			"$eq": nil,
		},
	})
	if err != nil {
		return false, err
	}
	return refreshToken != nil, nil
}
