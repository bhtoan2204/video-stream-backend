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

	bus.RegisterHandler("SearchUserQuery", func(q Query) (interface{}, error) {
		return searchUserHandler.Handle(q.(*query.SearchUserQuery))
	})

	// Register other command handlers here

	return bus
}
