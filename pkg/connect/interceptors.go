package connect

import (
	"context"

	"github.com/bufbuild/connect-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	grpc_codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			fullMethod := req.Spec().Procedure // e.g., "/acme.foo.v1.FooService/Bar"

			tracer := otel.Tracer("")

			ctx = extract(ctx, otel.GetTextMapPropagator(), req)

			name, attr := spanInfo(fullMethod, peerFromCtx(ctx))

			// Create a new span
			ctx, span := tracer.Start(
				trace.ContextWithRemoteSpanContext(ctx, trace.SpanContextFromContext(ctx)),
				name,
				trace.WithSpanKind(trace.SpanKindServer),
				trace.WithAttributes(attr...),
			)
			defer span.End()

			resp, err := next(ctx, req)
			if err != nil {
				s, _ := status.FromError(err)
				span.SetStatus(codes.Error, s.Message())
				span.SetAttributes(statusCodeAttr(s.Code()))
			} else {
				span.SetAttributes(statusCodeAttr(grpc_codes.OK))
			}

			return resp, err
		})
	}

	return interceptor
}
