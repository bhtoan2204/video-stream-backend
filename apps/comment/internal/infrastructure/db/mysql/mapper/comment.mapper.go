package mapper

import (
	"github.com/bhtoan2204/comment/internal/domain/entities"
	"github.com/bhtoan2204/comment/internal/infrastructure/db/mysql/model"
)

func CommentModelToEntity(comment *model.Comment) *entities.Comment {
	return &entities.Comment{
		AbstractModel: entities.AbstractModel{
			ID:        comment.ID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			DeletedAt: deletedAtToTimePointer(comment.DeletedAt),
		},
		VideoID:    comment.VideoID,
		UserID:     comment.UserID,
		Content:    comment.Content,
		ParentID:   comment.ParentID,
		LikeCount:  comment.LikeCount,
		ReplyCount: comment.ReplyCount,
		Status:     comment.Status,
	}
}

func CommentEntityToModel(comment *entities.Comment) *model.Comment {
	return &model.Comment{
		AbstractModel: model.AbstractModel{
			ID:        comment.ID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			DeletedAt: timePointerToDeletedAt(comment.DeletedAt),
		},
		VideoID:    comment.VideoID,
		UserID:     comment.UserID,
		Content:    comment.Content,
		ParentID:   comment.ParentID,
		LikeCount:  comment.LikeCount,
		ReplyCount: comment.ReplyCount,
		Status:     comment.Status,
	}
}
