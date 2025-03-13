package command

import (
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
	return validate.Struct(c)
}
