package repository

import (
	"errors"

	"github.com/bhtoan2204/user/internal/application/query"
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/domain/repository"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user *entities.User) (*entities.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *GormUserRepository) CreateUser(user *entities.User) (*entities.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *GormUserRepository) FindByQuery(q query.QueryOptions) ([]entities.User, error) {
	var users []entities.User
	dbQuery := r.db

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

	if err := dbQuery.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func (r *GormUserRepository) FindOneByQuery(q query.QueryOptions) (*entities.User, error) {
	var user entities.User
	dbQuery := r.db

	for key, value := range q.Filters {
		dbQuery = dbQuery.Where(key+" = ?", value)
	}

	if err := dbQuery.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
