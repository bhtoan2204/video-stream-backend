package repository

import (
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

// Create implements repository.UserRepository.
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

func (r *GormUserRepository) FindByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
