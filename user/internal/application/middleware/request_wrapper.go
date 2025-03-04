package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

func NewInstrumentedHandler(counter metric.Int64Counter, commonLabels []attribute.KeyValue) func(gin.HandlerFunc) gin.HandlerFunc {
	return func(handler gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			ctx := c.Request.Context()
			counter.Add(ctx, 1, metric.WithAttributes(commonLabels...))

			span := trace.SpanFromContext(ctx)
			bag := baggage.FromContext(ctx)
			var baggageAttributes []attribute.KeyValue
			baggageAttributes = append(baggageAttributes, commonLabels...)
			for _, member := range bag.Members() {
				baggageAttributes = append(baggageAttributes, attribute.String("baggage key:"+member.Key(), member.Value()))
			}
			span.SetAttributes(baggageAttributes...)

			handler(c)

			span.SetAttributes(commonLabels...)
		}
	}
}
