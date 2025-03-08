package mapper

import (
	"github.com/bhtoan2204/video/internal/domain/entities"
	"github.com/bhtoan2204/video/internal/infrastructure/db/mysql/model"
)

func VideoEntityToModel(video *entities.Video) *model.Video {
	return &model.Video{
		AbstractModel: model.AbstractModel{
			ID:        video.ID,
			CreatedAt: video.CreatedAt,
			UpdatedAt: video.UpdatedAt,
			DeletedAt: timePointerToDeletedAt(video.DeletedAt),
		},
		Title:        video.Title,
		Description:  video.Description,
		IsSearchable: video.IsSearchable,
		IsPublic:     video.IsPublic,
		VideoURL:     video.VideoURL,
		Bucket:       video.Bucket,
		ObjectKey:    video.ObjectKey,
		UploadedUser: video.UploadedUser,
	}
}

func VideoModelToEntity(video *model.Video) *entities.Video {
	return &entities.Video{
		AbstractModel: entities.AbstractModel{
			ID:        video.ID,
			CreatedAt: video.CreatedAt,
			UpdatedAt: video.UpdatedAt,
			DeletedAt: deletedAtToTimePointer(video.DeletedAt),
		},
		Title:        video.Title,
		Description:  video.Description,
		IsSearchable: video.IsSearchable,
		IsPublic:     video.IsPublic,
		VideoURL:     video.VideoURL,
		Bucket:       video.Bucket,
		ObjectKey:    video.ObjectKey,
		UploadedUser: video.UploadedUser,
	}
}
