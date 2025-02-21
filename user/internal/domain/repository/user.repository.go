package repository

import "github.com/bhtoan2204/user/internal/domain/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
}
