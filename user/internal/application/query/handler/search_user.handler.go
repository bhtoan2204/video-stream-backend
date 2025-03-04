package handler

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/query/query"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
)

type SearchUserQueryHandler struct {
	userService interfaces.UserServiceInterface
}

func NewSearchUserQueryHandler(userService interfaces.UserServiceInterface) *SearchUserQueryHandler {
	return &SearchUserQueryHandler{userService: userService}
}

func (h *SearchUserQueryHandler) Handle(ctx context.Context, q *query.SearchUserQuery) (*query.SearchUserQueryResult, error) {
	return h.userService.SearchUser(ctx, q)
}
