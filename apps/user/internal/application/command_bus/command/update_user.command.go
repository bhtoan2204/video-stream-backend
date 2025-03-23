package command

import "github.com/go-playground/validator/v10"

type UpdateUserCommand struct {
	ID        string `json:"id" validate:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	BirthDate string `json:"birth_date"`
	Avatar    string `json:"avatar"`
}

type UpdateUserCommandResult struct {
	Success bool `json:"success"`
}

func (c *UpdateUserCommand) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *UpdateUserCommand) CommandName() string {
	return "UpdateUserCommand"
}
