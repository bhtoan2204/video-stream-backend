package repository_interface

import (
	"context"

	"github.com/bhtoan2204/video/internal/domain/entities"
	"github.com/bhtoan2204/video/utils"
)

type VideoRepositoryInterface interface {
	CreateOne(context.Context, *entities.Video) (*entities.Video, error)
	UpdateOne(context.Context, *utils.QueryOptions, *entities.Video) (*entities.Video, error)
	DeleteOne(context.Context, *utils.QueryOptions) error
	FindOne(context.Context, *utils.QueryOptions) (*entities.Video, error)
	FindAll(context.Context, *utils.QueryOptions) ([]entities.Video, error)
}
