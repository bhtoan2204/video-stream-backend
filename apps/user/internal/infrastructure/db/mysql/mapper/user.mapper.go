package mapper

import (
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/model"
)

// Convert model.User to entities.User
func UserModelToEntity(user model.User) entities.User {
	roles := make([]*entities.Role, 0, len(user.Roles))
	for _, r := range user.Roles {
		if r != nil {
			roles = append(roles, RoleModelToEntity(*r))
		}
	}
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
		Roles:        roles,
	}
}

func UserEntityToModel(user entities.User) model.User {
	roles := make([]*model.Role, 0, len(user.Roles))
	for _, r := range user.Roles {
		if r != nil {
			roles = append(roles, RoleEntityToModel(*r))
		}
	}
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
		Roles:        roles,
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

func PermissionModelToEntity(perm model.Permission) *entities.Permission {
	return &entities.Permission{
		AbstractModel: entities.AbstractModel{
			ID:        perm.ID,
			CreatedAt: perm.CreatedAt,
			UpdatedAt: perm.UpdatedAt,
			DeletedAt: deletedAtToTimePointer(perm.DeletedAt),
		},
		Name:        perm.Name,
		Description: perm.Description,
	}
}

func PermissionEntityToModel(perm entities.Permission) *model.Permission {
	return &model.Permission{
		AbstractModel: model.AbstractModel{
			ID:        perm.ID,
			CreatedAt: perm.CreatedAt,
			UpdatedAt: perm.UpdatedAt,
			DeletedAt: timePointerToDeletedAt(perm.DeletedAt),
		},
		Name:        perm.Name,
		Description: perm.Description,
	}
}

func RoleModelToEntity(role model.Role) *entities.Role {
	permissions := make([]*entities.Permission, 0, len(role.Permissions))
	for _, perm := range role.Permissions {
		if perm != nil {
			permissions = append(permissions, PermissionModelToEntity(*perm))
		}
	}
	return &entities.Role{
		AbstractModel: entities.AbstractModel{
			ID:        role.ID,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
			DeletedAt: deletedAtToTimePointer(role.DeletedAt),
		},
		Name:        role.Name,
		Permissions: permissions,
	}
}

func RoleEntityToModel(role entities.Role) *model.Role {
	permissions := make([]*model.Permission, 0, len(role.Permissions))
	for _, perm := range role.Permissions {
		if perm != nil {
			permissions = append(permissions, PermissionEntityToModel(*perm))
		}
	}
	return &model.Role{
		AbstractModel: model.AbstractModel{
			ID:        role.ID,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
			DeletedAt: timePointerToDeletedAt(role.DeletedAt),
		},
		Name:        role.Name,
		Permissions: permissions,
	}
}
