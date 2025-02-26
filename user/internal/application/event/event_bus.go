package event

import (
	"errors"

	"github.com/bhtoan2204/user/global"
	"go.uber.org/zap"
)

type Event interface {
	EventName() string
}

type HandlerFunc func(Event) (interface{}, error)

type EventBus struct {
	handlers map[string]HandlerFunc
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string]HandlerFunc),
	}
}

func (bus *EventBus) RegisterHandler(commandName string, handler HandlerFunc) {
	bus.handlers[commandName] = handler
}

func (bus *EventBus) Dispatch(cmd Event) (interface{}, error) {
	handler, exists := bus.handlers[cmd.EventName()]
	if !exists {
		global.Logger.Error("No handler registered for command", zap.String("command", cmd.EventName()))
		return nil, errors.New("no handler registered for command: " + cmd.EventName())
	}
	return handler(cmd)
}

func (bus *EventBus) DispatchAsync(cmd Event) {
	go func() {
		if _, err := bus.Dispatch(cmd); err != nil {
			global.Logger.Error("Failed to dispatch event", zap.Error(err))
		}
	}()
}
