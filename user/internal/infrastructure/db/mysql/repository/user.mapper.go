package repository

import "github.com/bhtoan2204/user/internal/domain/entities"

func toDBUser(user *entities.User) *entities.User {
	return user
}

func toDomainUser(user *entities.User) *entities.User {
	return user
}
