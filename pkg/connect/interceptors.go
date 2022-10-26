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

// TODO this can go away eventually when we get an official interceptor:
// https://github.com/bufbuild/connect-go/issues/344
// For now, we're inspired by:
// https://github.com/open-telemetry/opentelemetry-go-contrib/blob/instrumentation/google.golang.org/grpc/otelgrpc/v0.36.4/instrumentation/google.golang.org/grpc/otelgrpc/interceptor.go#L307
func connectInterceptorForSpan() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			fullMethod := req.Spec().Procedure // e.g., "/acme.foo.v1.FooService/Bar"

			tr := otel.Tracer("")

			ctx = extract(ctx, otel.GetTextMapPropagator(), req)

			name, attr := spanInfo(fullMethod, peerFromCtx(ctx))

			// Create a new span
			ctx, span := tr.Start(
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
	return connect.UnaryInterceptorFunc(interceptor)
}
