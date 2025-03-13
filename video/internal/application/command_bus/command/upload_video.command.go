package command

import (
	"github.com/bhtoan2204/video/internal/domain/entities"
	"github.com/go-playground/validator"
)

type UploadVideoCommand struct {
	FileKey      string `json:"file_key" validate:"required"`
	FileName     string `json:"file_name" validate:"required"`
	ContentType  string `json:"content_type" validate:"required"`
	FileSize     int64  `json:"file_size" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description"`
	IsPublic     bool   `json:"is_public"`
	IsSearchable bool   `json:"is_searchable"`
	UploadedUser string `json:"uploaded_user" validate:"required"`
}

type UploadVideoCommandResult struct {
	Result *entities.Video `json:"result"`
}

func (*UploadVideoCommand) CommandName() string {
	return "UploadVideoCommand"
}

func (c *UploadVideoCommand) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
