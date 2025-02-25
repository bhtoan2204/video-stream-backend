package mapper

import (
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/model"
)

// Convert model.User to entities.User
func UserModelToEntity(user model.User) entities.User {
	return entities.User{
		AbstractModel: entities.AbstractModel{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: deletedAtToTimePointer(user.DeletedAt),
		},
		Username:     user.Username,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Phone:        user.Phone,
		BirthDate:    user.BirthDate,
		Address:      user.Address,
		PasswordHash: user.PasswordHash,
		PinCode:      user.PinCode,
		Status:       entities.Status(user.Status),
	}
}

func UserEntityToModel(user entities.User) model.User {
	return model.User{
		AbstractModel: model.AbstractModel{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: timePointerToDeletedAt(user.DeletedAt),
		},
		Username:     user.Username,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Phone:        user.Phone,
		BirthDate:    user.BirthDate,
		Address:      user.Address,
		PasswordHash: user.PasswordHash,
		PinCode:      user.PinCode,
		Status:       model.Status(user.Status),
	}
}

func UserModelsToEntities(users []model.User) []entities.User {
	entitiesList := make([]entities.User, len(users))
	for i, user := range users {
		entitiesList[i] = UserModelToEntity(user)
	}
	return entitiesList
}

func UserEntitiesToModels(users []entities.User) []model.User {
	modelsList := make([]model.User, len(users))
	for i, user := range users {
		modelsList[i] = UserEntityToModel(user)
	}
	return modelsList
}
