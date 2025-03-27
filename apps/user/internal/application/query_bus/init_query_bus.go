package query_bus

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/query_bus/handler"
	"github.com/bhtoan2204/user/internal/application/query_bus/query"
	"github.com/bhtoan2204/user/internal/application/shared"
)

func SetUpQueryBus(deps *shared.ServiceDependencies) *QueryBus {
	bus := NewQueryBus()

	// Register User command handlers
	searchUserHandler := handler.NewSearchUserQueryHandler(deps.UserService)
	getUserProfileHandler := handler.NewGetUserProfileQueryHandler(deps.UserService)

	bus.RegisterHandler("SearchUserQuery", func(ctx context.Context, q Query) (interface{}, error) {
		return searchUserHandler.Handle(ctx, q.(*query.SearchUserQuery))
	})

	bus.RegisterHandler("GetUserProfileQuery", func(ctx context.Context, q Query) (interface{}, error) {
		return getUserProfileHandler.Handle(ctx, q.(*query.GetUserProfileQuery))
	})

	// Register other command handlers here

	return bus
}
