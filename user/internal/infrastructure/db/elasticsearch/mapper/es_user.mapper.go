package mapper

import (
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/infrastructure/db/elasticsearch/model"
)

func ESUserModelToEntity(model model.ESUser) entities.User {
	return entities.User{
		AbstractModel: entities.AbstractModel{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
			DeletedAt: model.DeletedAt,
		},
		Username:     model.Username,
		Email:        model.Email,
		FirstName:    model.FirstName,
		LastName:     model.LastName,
		Phone:        model.Phone,
		BirthDate:    model.BirthDate,
		Address:      model.Address,
		PasswordHash: model.PasswordHash,
		PinCode:      model.PinCode,
		Status:       1,
	}
}

func ESUserEntityToModel(entity entities.User) model.ESUser {
	return model.ESUser{
		ID:           entity.ID,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
		DeletedAt:    entity.DeletedAt,
		Username:     entity.Username,
		Email:        entity.Email,
		FirstName:    entity.FirstName,
		LastName:     entity.LastName,
		Phone:        entity.Phone,
		BirthDate:    entity.BirthDate,
		Address:      entity.Address,
		PasswordHash: entity.PasswordHash,
		PinCode:      entity.PinCode,
	}
}
