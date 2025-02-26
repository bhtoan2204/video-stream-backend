package event

import (
	"github.com/bhtoan2204/user/internal/application/event/event"
	"github.com/bhtoan2204/user/internal/application/event/handler"
	"github.com/bhtoan2204/user/internal/application/shared"
)

func SetUpEventBus(deps *shared.ListenerDependencies) *EventBus {
	bus := NewEventBus()

	indexUserHandler := handler.NewIndexUserEventHandler(deps.UserListener)
	//dbserver1.user.users
	bus.RegisterHandler("IndexUserEvent", func(e Event) (interface{}, error) {
		return indexUserHandler.Handle(e.(*event.IndexUserEvent))
	})

	return bus
}
