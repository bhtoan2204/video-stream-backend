package query

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
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
	queryName := query.QueryName()

	tracer := otel.Tracer("query-bus")
	ctx, commandSpan := tracer.Start(ctx, queryName)
	defer commandSpan.End()

	handler, exists := bus.handlers[queryName]
	if !exists {
		err := errors.New("no handler registered for event: " + queryName)
		commandSpan.RecordError(err)
		return nil, err
	}

	result, err := handler(ctx, query)
	if err != nil {
		commandSpan.RecordError(err)
	}
	return result, err
}
