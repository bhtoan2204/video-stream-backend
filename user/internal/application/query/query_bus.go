package query

import (
	"context"
	"errors"
)

type Query interface {
	QueryName() string
}

type HandlerFunc func(context.Context, Query) (interface{}, error)

type QueryBus struct {
	handlers map[string]HandlerFunc
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[string]HandlerFunc),
	}
}

func (bus *QueryBus) RegisterHandler(queryName string, handler HandlerFunc) {
	bus.handlers[queryName] = handler
}

func (bus *QueryBus) Dispatch(ctx context.Context, query Query) (interface{}, error) {
	handler, exists := bus.handlers[query.QueryName()]
	if !exists {
		return nil, errors.New("no handler registered for query: " + query.QueryName())
	}
	return handler(ctx, query)
}
