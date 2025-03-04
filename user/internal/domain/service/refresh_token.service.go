package service

import (
	"context"
	"time"

	"github.com/bhtoan2204/user/internal/domain/entities"
	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
)

type RefreshTokenService struct {
	refreshTokenRepository repository.RefreshTokenRepositoryInterface
}

func NewRefreshTokenService(refreshTokenRepository repository.RefreshTokenRepositoryInterface) *RefreshTokenService {
	return &RefreshTokenService{
		refreshTokenRepository: refreshTokenRepository,
	}
}

func (s *RefreshTokenService) CreateRefreshToken(ctx context.Context, token string, userId string, expires_at time.Time) error {
	return s.refreshTokenRepository.Create(ctx, &entities.RefreshToken{
		Token:     token,
		UserID:    userId,
		ExpiresAt: expires_at,
	})
}

func (s *RefreshTokenService) HardDeleteByQuery(ctx context.Context, query map[string]interface{}) error {
	return s.refreshTokenRepository.HardDeleteByQuery(ctx, &query)
}

func (s *RefreshTokenService) RevokedByQuery(ctx context.Context, query map[string]interface{}) error {
	return s.refreshTokenRepository.UpdateByQuery(ctx, &query, &map[string]interface{}{
		"revoked_at": time.Now(),
	})
}

func (s *RefreshTokenService) DeleteByQuery(ctx context.Context, query map[string]interface{}) error {
	return s.refreshTokenRepository.DeleteByQuery(ctx, &query)
}

func (s *RefreshTokenService) CheckRefreshToken(ctx context.Context, token string) (bool, error) {
	refreshToken, err := s.refreshTokenRepository.FindOneByQuery(ctx, &map[string]interface{}{
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
