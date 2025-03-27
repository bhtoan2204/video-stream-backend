package command_bus

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
)

type Command interface {
	CommandName() string
}

type HandlerFunc func(context.Context, Command) (interface{}, error)

type CommandBus struct {
	handlers map[string]HandlerFunc
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]HandlerFunc),
	}
}

func (bus *CommandBus) RegisterHandler(commandName string, handler HandlerFunc) {
	bus.handlers[commandName] = handler
}

func (bus *CommandBus) Dispatch(ctx context.Context, cmd Command) (interface{}, error) {
	commandName := cmd.CommandName()

	tracer := otel.Tracer("command-bus")
	ctx, commandSpan := tracer.Start(ctx, commandName)
	defer commandSpan.End()

	handler, exists := bus.handlers[commandName]
	if !exists {
		err := errors.New("no handler registered for command: " + commandName)
		commandSpan.RecordError(err)
		return nil, err
	}

	result, err := handler(ctx, cmd)
	if err != nil {
		commandSpan.RecordError(err)
	}
	return result, err
}
