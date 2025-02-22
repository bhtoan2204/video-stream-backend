package query

import common "github.com/bhtoan2204/user/internal/application/common/query"

type SearchUserQuery struct {
	Query string `form:"query" binding:"required"`
}

type SearchUserQueryResult struct {
	Result *[]common.UserResult `json:"result"`
}

func (q SearchUserQuery) QueryName() string {
	return "SearchUserQuery"
}
