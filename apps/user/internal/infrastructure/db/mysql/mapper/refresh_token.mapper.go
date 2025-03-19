package mapper

import (
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/model"
)

func RefreshTokenModelToEntity(refresh_token model.RefreshToken) entities.RefreshToken {
	return entities.RefreshToken{
		AbstractModel: entities.AbstractModel{
			ID:        refresh_token.ID,
			CreatedAt: refresh_token.CreatedAt,
			UpdatedAt: refresh_token.UpdatedAt,
			DeletedAt: deletedAtToTimePointer(refresh_token.DeletedAt),
		},
		UserID:    refresh_token.UserID,
		Token:     refresh_token.Token,
		ExpiresAt: refresh_token.ExpiresAt,
	}
}

func RefreshTokenEntityToModel(refresh_token entities.RefreshToken) model.RefreshToken {
	return model.RefreshToken{
		AbstractModel: model.AbstractModel{
			ID:        refresh_token.ID,
			CreatedAt: refresh_token.CreatedAt,
			UpdatedAt: refresh_token.UpdatedAt,
			DeletedAt: timePointerToDeletedAt(refresh_token.DeletedAt),
		},
		UserID:    refresh_token.UserID,
		Token:     refresh_token.Token,
		ExpiresAt: refresh_token.ExpiresAt,
	}
}
