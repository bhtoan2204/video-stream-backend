package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UploadVideoRequest struct {
	FileKey      string `json:"file_key" validate:"required"`
	FileName     string `json:"file_name" validate:"required"`
	ContentType  string `json:"content_type" validate:"required"`
	FileSize     int64  `json:"file_size" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description"`
	IsPublic     bool   `json:"is_public"`
	IsSearchable bool   `json:"is_searchable"`
	// UploadedUser string `json:"uploaded_user" validate:"required"`
}

func (c *UploadVideoRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		var errorMessage string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage += fmt.Sprintf("Field: %s, Error: %s, Value: %v\n",
				err.Field(),
				err.Tag(),
				err.Value())
		}
		return fmt.Errorf(errorMessage)
	}
	return nil
}
