package command

import (
	"fmt"

	"github.com/bhtoan2204/video/internal/domain/entities"
	"github.com/go-playground/validator"
)

type GetVideoByURLCommand struct {
	URL string `json:"url" validate:"required"`
}

type GetVideoByURLCommandResult struct {
	Result *entities.Video `json:"result"`
}

func (*GetVideoByURLCommand) CommandName() string {
	return "GetVideoByURLCommand"
}

func (c *GetVideoByURLCommand) Validate() error {
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
