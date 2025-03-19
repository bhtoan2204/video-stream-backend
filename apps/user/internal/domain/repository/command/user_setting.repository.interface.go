package repository_interface

import "context"

type UserSettingRepositoryInterface interface {
	UpdateByUserId(ctx context.Context, userID *string, update *map[string]interface{}) error
}
