package dto

import "github.com/go-playground/validator/v10"

type GetProfileRequest struct {
}

func (c *GetProfileRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
