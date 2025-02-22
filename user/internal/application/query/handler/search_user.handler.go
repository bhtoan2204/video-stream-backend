package handler

import (
	"github.com/bhtoan2204/user/internal/application/interfaces"
	"github.com/bhtoan2204/user/internal/application/query/query"
)

type SearchUserQueryHandler struct {
	userService interfaces.UserServiceInterface
}

func NewSearchUserQueryHandler(userService interfaces.UserServiceInterface) *SearchUserQueryHandler {
	return &SearchUserQueryHandler{userService: userService}
}

func (h *SearchUserQueryHandler) Handle(q *query.SearchUserQuery) (*query.SearchUserQueryResult, error) {
	return h.userService.SearchUser(q)
}
