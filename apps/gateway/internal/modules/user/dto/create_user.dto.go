package dto

import "github.com/go-playground/validator/v10"

type CreateUserRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=20"`
	Password  string `json:"password" validate:"required,min=8"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
	Address   string `json:"address" validate:"required"`
}

func (c *CreateUserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
