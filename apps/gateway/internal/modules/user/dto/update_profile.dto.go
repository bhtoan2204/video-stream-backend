package dto

import "github.com/go-playground/validator/v10"

type UpdateProfileRequest struct {
	ID        string `json:"id" validate:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	BirthDate string `json:"birth_date"`
	Avatar    string `json:"avatar"`
}

func (c *UpdateProfileRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
