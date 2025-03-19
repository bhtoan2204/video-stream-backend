package repository

import (
	"context"
	"errors"
	"time"

	"github.com/bhtoan2204/video/internal/domain/entities"
	repository_interface "github.com/bhtoan2204/video/internal/domain/repository/command"
	"github.com/bhtoan2204/video/internal/infrastructure/db/mysql/mapper"
	"github.com/bhtoan2204/video/utils"
	"gorm.io/gorm"
)

type GormVideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) repository_interface.VideoRepositoryInterface {
	return &GormVideoRepository{db: db}
}

func (g *GormVideoRepository) CreateOne(ctx context.Context, video *entities.Video) (*entities.Video, error) {
	videoModel := mapper.VideoEntityToModel(video)
	err := g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		execution := g.db.Create(&videoModel)
		if execution.Error != nil {
			return execution.Error
		}
		if execution.RowsAffected == 0 {
			return errors.New("no rows affected")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	videoEntity := mapper.VideoModelToEntity(videoModel)
	return videoEntity, nil
}

func (g *GormVideoRepository) DeleteOne(ctx context.Context, q *utils.QueryOptions) error {
	// find and soft delete
	now := time.Now()
	dbQuery := BuildDbQuery(g.db, q)
	err := g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := dbQuery.First(&entities.Video{}).Update("deleted_at", now).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (g *GormVideoRepository) FindAll(ctx context.Context, q *utils.QueryOptions) ([]entities.Video, error) {
	dbQuery := BuildDbQuery(g.db, q)
	var videos []entities.Video
	err := dbQuery.WithContext(ctx).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (g *GormVideoRepository) FindOne(ctx context.Context, q *utils.QueryOptions) (*entities.Video, error) {
	dbQuery := BuildDbQuery(g.db, q)
	var video entities.Video
	err := dbQuery.WithContext(ctx).First(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func (g *GormVideoRepository) UpdateOne(ctx context.Context, q *utils.QueryOptions, video *entities.Video) (*entities.Video, error) {
	dbQuery := BuildDbQuery(g.db, q)
	err := g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := dbQuery.First(&entities.Video{}).Updates(video).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return video, nil
}
