package repository

import (
	"context"

	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/model"
	"gorm.io/gorm"
)

type GormUserSettingRepository struct {
	db *gorm.DB
}

func NewUserSettingRepository(db *gorm.DB) repository.UserSettingRepositoryInterface {
	return &GormUserSettingRepository{db: db}
}

func (g *GormUserSettingRepository) UpdateByUserId(ctx context.Context, userID *string, update *map[string]interface{}) error {
	if (*update)["user_id"] != nil {
		delete(*update, "user_id")
	}
	return g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(&model.UserSettings{}).
			Where("user_id = ?", userID).
			Updates(update).Error
	})
}
