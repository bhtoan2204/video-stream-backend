package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/bhtoan2204/comment/internal/domain/entities"
	"github.com/bhtoan2204/comment/internal/infrastructure/db/mysql/mapper"
	"github.com/bhtoan2204/comment/internal/infrastructure/db/mysql/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(ctx context.Context, comment *entities.Comment) (*entities.Comment, error) {
	model := mapper.CommentEntityToModel(comment)
	fmt.Printf("%+v\n", model)
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		execution := tx.Create(model)
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
	entity := mapper.CommentModelToEntity(model)
	return entity, nil
}

func (r *CommentRepository) Update(ctx context.Context, comment *entities.Comment) (*entities.Comment, error) {
	model := mapper.CommentEntityToModel(comment)
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		execution := tx.Save(model)
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
	entity := mapper.CommentModelToEntity(model)
	return entity, nil
}

func (r *CommentRepository) Delete(ctx context.Context, id string) error {
	// soft delete
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		execution := tx.Delete(&model.Comment{}, id)
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
