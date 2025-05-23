package repository

import (
	"context"
	"errors"

	"github.com/bhtoan2204/user/internal/domain/entities"
	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/mapper"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/model"
	"github.com/bhtoan2204/user/utils"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepositoryInterface {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	userModel := mapper.UserEntityToModel(*user)
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		execution := r.db.Create(&userModel)
		if execution.Error != nil {
			return execution.Error
		}
		if execution.RowsAffected == 0 {
			return errors.New("no rows affected")
		}
		var adminRole model.Role
		if err := tx.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
			return err
		}
		if err := tx.Model(&userModel).Association("Roles").Append(&adminRole); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	userEntity := mapper.UserModelToEntity(userModel)
	return &userEntity, nil
}

func (r *GormUserRepository) FindByQuery(ctx context.Context, q *utils.QueryOptions) ([]entities.User, error) {
	userListModel := mapper.UserEntitiesToModels([]entities.User{})
	dbQuery := r.db.WithContext(ctx)
	for key, value := range q.Filters {
		dbQuery = dbQuery.Where(key+" = ?", value)
	}
	if q.OrderBy != "" {
		dbQuery = dbQuery.Order(q.OrderBy)
	}
	if q.Limit > 0 {
		dbQuery = dbQuery.Limit(q.Limit)
	}
	if q.Offset > 0 {
		dbQuery = dbQuery.Offset(q.Offset)
	}
	if err := dbQuery.Model(&model.User{}).Find(&userListModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.UserModelsToEntities(userListModel), nil
}

func (r *GormUserRepository) FindOneByQuery(ctx context.Context, q *utils.QueryOptions) (*entities.User, error) {
	var userModel model.User
	dbQuery := r.db.WithContext(ctx)

	for key, value := range q.Filters {
		dbQuery = dbQuery.Where(key+" = ?", value)
	}

	// Preload roles and permissions
	dbQuery = dbQuery.Preload("Roles").Preload("Roles.Permissions")

	if err := dbQuery.Model(&model.User{}).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	userEntity := mapper.UserModelToEntity(userModel)
	return &userEntity, nil
}

func (r *GormUserRepository) UpdateOne(ctx context.Context, user *entities.User) error {
	userModel := mapper.UserEntityToModel(*user)
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", user.ID).Updates(userModel).Error
}
