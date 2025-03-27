package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreateCommentRequest struct {
	VideoId  string `json:"video_id" validate:"required"`
	Content  string `json:"content" validate:"required,min=1"`
	ParentID string `json:"parent_id,omitempty"`
}

func (c *CreateCommentRequest) Validate() error {
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
