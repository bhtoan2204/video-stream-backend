package command

import (
	"fmt"

	"github.com/bhtoan2204/comment/internal/domain/entities"
	"github.com/go-playground/validator"
)

type CreateCommentCommand struct {
	VideoId  string `json:"video_id" validate:"required"`
	UserId   string `json:"user_id" validate:"required"`
	Content  string `json:"content" validate:"required,min=1"` // min size 1
	ParentID string `json:"parent_id"`
}

type CreateCommentCommandResult struct {
	Result *entities.Comment `json:"result"`
}

func (*CreateCommentCommand) CommandName() string {
	return "CreateCommentCommand"
}

func (c *CreateCommentCommand) Validate() error {
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
