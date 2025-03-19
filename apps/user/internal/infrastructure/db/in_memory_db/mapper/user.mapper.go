package mapper_test

import (
	"github.com/bhtoan2204/user/internal/domain/entities"
	model_test "github.com/bhtoan2204/user/internal/infrastructure/db/in_memory_db/model"
)

// Convert model_test.User to entities.User
func UserModelToEntity(user model_test.User) entities.User {
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

func UserEntityToModel(user entities.User) model_test.User {
	roles := make([]*model_test.Role, 0, len(user.Roles))
	for _, r := range user.Roles {
		if r != nil {
			roles = append(roles, RoleEntityToModel(*r))
		}
	}
	return model_test.User{
		AbstractModel: model_test.AbstractModel{
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
		Status:       model_test.Status(user.Status),
		Roles:        roles,
	}
}

func UserModelsToEntities(users []model_test.User) []entities.User {
	entitiesList := make([]entities.User, len(users))
	for i, user := range users {
		entitiesList[i] = UserModelToEntity(user)
	}
	return entitiesList
}

func UserEntitiesToModels(users []entities.User) []model_test.User {
	modelsList := make([]model_test.User, len(users))
	for i, user := range users {
		modelsList[i] = UserEntityToModel(user)
	}
	return modelsList
}

func PermissionModelToEntity(perm model_test.Permission) *entities.Permission {
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

func PermissionEntityToModel(perm entities.Permission) *model_test.Permission {
	return &model_test.Permission{
		AbstractModel: model_test.AbstractModel{
			ID:        perm.ID,
			CreatedAt: perm.CreatedAt,
			UpdatedAt: perm.UpdatedAt,
			DeletedAt: timePointerToDeletedAt(perm.DeletedAt),
		},
		Name:        perm.Name,
		Description: perm.Description,
	}
}

func RoleModelToEntity(role model_test.Role) *entities.Role {
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

func RoleEntityToModel(role entities.Role) *model_test.Role {
	permissions := make([]*model_test.Permission, 0, len(role.Permissions))
	for _, perm := range role.Permissions {
		if perm != nil {
			permissions = append(permissions, PermissionEntityToModel(*perm))
		}
	}
	return &model_test.Role{
		AbstractModel: model_test.AbstractModel{
			ID:        role.ID,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
			DeletedAt: timePointerToDeletedAt(role.DeletedAt),
		},
		Name:        role.Name,
		Permissions: permissions,
	}
}
