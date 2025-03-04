package command

import (
	"context"
	"errors"
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
	handler, exists := bus.handlers[cmd.CommandName()]
	if !exists {
		return nil, errors.New("no handler registered for command: " + cmd.CommandName())
	}
	return handler(ctx, cmd)
}
