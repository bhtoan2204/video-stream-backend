package dto

import "github.com/go-playground/validator/v10"

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
	return validate.Struct(c)
}
