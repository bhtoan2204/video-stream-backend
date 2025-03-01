package handler

import (
	"github.com/bhtoan2204/user/internal/application/query/query"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
)

type GetUserProfileQueryHandler struct {
	userService interfaces.UserServiceInterface
}

func NewGetUserProfileQueryHandler(userService interfaces.UserServiceInterface) *GetUserProfileQueryHandler {
	return &GetUserProfileQueryHandler{userService: userService}
}

func (h *GetUserProfileQueryHandler) Handle(q *query.GetUserProfileQuery) (*query.GetUserProfileQueryResult, error) {
	return h.userService.GetUserProfile(q)
}
