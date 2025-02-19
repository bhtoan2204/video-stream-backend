package repository

import "github.com/bhtoan2204/user/internal/domain/model"

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindById(id uint) (*model.User, error)
}
