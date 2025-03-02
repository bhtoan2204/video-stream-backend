package query

import (
	"github.com/bhtoan2204/user/internal/application/query/handler"
	"github.com/bhtoan2204/user/internal/application/query/query"
	"github.com/bhtoan2204/user/internal/application/shared"
)

func SetUpQueryBus(deps *shared.ServiceDependencies) *QueryBus {
	bus := NewQueryBus()

	// Register User command handlers
	searchUserHandler := handler.NewSearchUserQueryHandler(deps.UserService)
	getUserProfileHandler := handler.NewGetUserProfileQueryHandler(deps.UserService)

	bus.RegisterHandler("SearchUserQuery", func(q Query) (interface{}, error) {
		return searchUserHandler.Handle(q.(*query.SearchUserQuery))
	})

	bus.RegisterHandler("GetUserProfileQuery", func(q Query) (interface{}, error) {
		return getUserProfileHandler.Handle(q.(*query.GetUserProfileQuery))
	})

	// Register other command handlers here

	return bus
}
