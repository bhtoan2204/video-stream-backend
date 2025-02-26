package repository

import (
	"errors"
	"time"

	"github.com/bhtoan2204/user/internal/domain/entities"
	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/mapper"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/model"
	"gorm.io/gorm"
)

type GormRefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) repository.RefreshTokenRepository {
	return &GormRefreshTokenRepository{db: db}
}

// Create implements command.RefreshTokenRepository.
func (r *GormRefreshTokenRepository) Create(refreshToken *entities.RefreshToken) error {
	refreshTokenModel := model.RefreshToken{
		UserID:    refreshToken.UserID,
		Token:     refreshToken.Token,
		ExpiresAt: refreshToken.ExpiresAt,
	}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		execution := r.db.Create(&refreshTokenModel)
		if execution.Error != nil {
			return execution.Error
		}
		if execution.RowsAffected == 0 {
			return errors.New("no rows affected")
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteByQuery implements command.RefreshTokenRepository.
func (g *GormRefreshTokenRepository) DeleteByQuery(query map[string]interface{}) error {
	now := time.Now()
	err := g.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.RefreshToken{}).Where(query).Update("deleted_at", &now).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// FindOneByQuery implements command.RefreshTokenRepository.
func (g *GormRefreshTokenRepository) FindOneByQuery(query map[string]interface{}) (*entities.RefreshToken, error) {
	refreshTokenModel := model.RefreshToken{}
	dbQuery := BuildQuery(g.db, query)
	if err := dbQuery.First(&refreshTokenModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	refreshTokenEntity := mapper.RefreshTokenModelToEntity(refreshTokenModel)
	return &refreshTokenEntity, nil
}

// UpdateByQuery implements command.RefreshTokenRepository.
func (g *GormRefreshTokenRepository) UpdateByQuery(query map[string]interface{}, update map[string]interface{}) error {
	err := g.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.RefreshToken{}).Where(query).Updates(update).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *GormRefreshTokenRepository) HardDeleteByQuery(query map[string]interface{}) error {
	err := g.db.Transaction(func(tx *gorm.DB) error {
		if err := g.db.Where(query).Delete(&model.RefreshToken{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
