package event

import (
	"context"
	"errors"

	"github.com/bhtoan2204/user/global"
	"go.uber.org/zap"
)

type Event interface {
	EventName() string
}

type HandlerFunc func(context.Context, Event) (interface{}, error)

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

func (bus *EventBus) Dispatch(ctx context.Context, cmd Event) (interface{}, error) {
	handler, exists := bus.handlers[cmd.EventName()]
	if !exists {
		global.Logger.Error("No handler registered for command", zap.String("command", cmd.EventName()))
		return nil, errors.New("no handler registered for command: " + cmd.EventName())
	}
	return handler(ctx, cmd)
}

func (bus *EventBus) DispatchAsync(ctx context.Context, cmd Event) {
	go func() {
		if _, err := bus.Dispatch(ctx, cmd); err != nil {
			global.Logger.Error("Failed to dispatch event", zap.Error(err))
		}
	}()
}
