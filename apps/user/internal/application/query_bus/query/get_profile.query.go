package query

import common "github.com/bhtoan2204/user/internal/application/common/query"

type GetUserProfileQuery struct {
	ID string `form:"id"`
}

type GetUserProfileQueryResult struct {
	Result *common.UserResult `json:"result"`
}

func (q *GetUserProfileQuery) QueryName() string {
	return "GetUserProfileQuery"
}
