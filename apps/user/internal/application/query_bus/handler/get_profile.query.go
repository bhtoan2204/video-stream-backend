package handler

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/query_bus/query"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
)

type GetUserProfileQueryHandler struct {
	userService interfaces.UserServiceInterface
}

func NewGetUserProfileQueryHandler(userService interfaces.UserServiceInterface) *GetUserProfileQueryHandler {
	return &GetUserProfileQueryHandler{userService: userService}
}

func (h *GetUserProfileQueryHandler) Handle(ctx context.Context, q *query.GetUserProfileQuery) (*query.GetUserProfileQueryResult, error) {
	return h.userService.GetUserProfile(ctx, q)
}
