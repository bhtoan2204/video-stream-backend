package dto

import "github.com/go-playground/validator/v10"

type CreateCommentRequest struct {
	VideoId  string `json:"video_id" validate:"required"`
	UserId   string `json:"user_id" validate:"required"`
	Content  string `json:"content" validate:"required,min=1"`
	ParentID string `json:"parent_id,omitempty"`
}

func (c *CreateCommentRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
