package interceptor

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc"
)

func UnaryTracingInterceptor() grpc.UnaryServerInterceptor {
	tracer := otel.Tracer("grpc-server")
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		ctx, span := tracer.Start(ctx, info.FullMethod)
		defer span.End()

		resp, err = handler(ctx, req)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		} else {
			span.SetStatus(codes.Ok, "OK")
		}
		return resp, err
	}
}
