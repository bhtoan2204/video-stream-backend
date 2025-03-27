package event_bus

import (
	"context"
	"errors"

	"github.com/bhtoan2204/user/global"
	"go.opentelemetry.io/otel"
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

func (bus *EventBus) Dispatch(ctx context.Context, event Event) (interface{}, error) {
	eventName := event.EventName()

	tracer := otel.Tracer("event-bus")
	ctx, commandSpan := tracer.Start(ctx, eventName)
	defer commandSpan.End()

	handler, exists := bus.handlers[eventName]
	if !exists {
		err := errors.New("no handler registered for event: " + eventName)
		commandSpan.RecordError(err)
		return nil, err
	}

	result, err := handler(ctx, event)
	if err != nil {
		commandSpan.RecordError(err)
	}
	return result, err
}

func (bus *EventBus) DispatchAsync(ctx context.Context, cmd Event) {
	go func() {
		if _, err := bus.Dispatch(ctx, cmd); err != nil {
			global.Logger.Error("Failed to dispatch event", zap.Error(err))
		}
	}()
}
