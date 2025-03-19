package repository_interface

import (
	"context"

	"github.com/bhtoan2204/comment/internal/domain/entities"
)

type CommentRepositoryInterface interface {
	Create(context.Context, *entities.Comment) (*entities.Comment, error)
	Update(context.Context, *entities.Comment) (*entities.Comment, error)
	Delete(context.Context, string) error
}
