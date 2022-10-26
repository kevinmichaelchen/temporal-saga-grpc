package connect

import (
	"context"
	"github.com/bufbuild/connect-go"
	"go.opentelemetry.io/otel"
)

func UnaryInterceptors() []connect.Interceptor {
	return []connect.Interceptor{
		connectInterceptorForSpan(),
	}
}

func connectInterceptorForSpan() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			name := req.Spec().Procedure

			tr := otel.Tracer("")

			// Create a new span
			ctx, span := tr.Start(ctx, name)
			defer span.End()

			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
