package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type SearchUserRequest struct {
	Query         string `form:"query"`
	Page          int    `form:"page"`
	Limit         int    `form:"limit"`
	SortBy        string `form:"sort_by"`        // ("username", "email", …)
	SortDirection string `form:"sort_direction"` // ("asc" or "desc")
}

func (q *SearchUserRequest) SetDefaults() {
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.Limit <= 0 {
		q.Limit = 10 // default limit
	}
	if q.SortBy == "" {
		q.SortBy = "username" // default sort by username
	}
	if q.SortDirection == "" {
		q.SortDirection = "asc" // default ascending order
	}
}

func (c *SearchUserRequest) Validate() error {
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
