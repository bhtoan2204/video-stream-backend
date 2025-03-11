package query

import common "github.com/bhtoan2204/user/internal/application/common/query"

type SearchUserQuery struct {
	Query         string `form:"query"`
	Page          int    `form:"page"`
	Limit         int    `form:"limit"`
	SortBy        string `form:"sort_by"`        // ("username", "email", …)
	SortDirection string `form:"sort_direction"` // ("asc" or "desc")
}

func (q *SearchUserQuery) SetDefaults() {
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

type SearchUserQueryResult struct {
	Result         *[]common.UserResult `json:"result"`
	PaginateResult *PaginateResult      `json:"paginate_result"`
}

func (q SearchUserQuery) QueryName() string {
	return "SearchUserQuery"
}
