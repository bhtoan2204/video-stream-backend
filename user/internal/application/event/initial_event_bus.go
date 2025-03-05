package event

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/event/event"
	"github.com/bhtoan2204/user/internal/application/event/handler"
	"github.com/bhtoan2204/user/internal/application/shared"
)

func SetUpEventBus(deps *shared.ListenerDependencies) *EventBus {
	bus := NewEventBus()

	indexUserHandler := handler.NewIndexUserEventHandler(deps.UserListener)
	//user_database.user.users
	bus.RegisterHandler("IndexUserEvent", func(ctx context.Context, e Event) (interface{}, error) {
		return indexUserHandler.Handle(ctx, e.(*event.IndexUserEvent))
	})

	return bus
}
