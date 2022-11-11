package connect

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
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

			logBaggage(ctx)

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

func logBaggage(ctx context.Context) {
	// EXTRACT BAGGAGE!!
	bg := baggage.FromContext(ctx)
	members := bg.Members()
	logrus.WithField("otel.baggage.num_members", len(members)).Info("OTel baggage")
	//for _, m := range bg.Members() {
	//	span.SetAttributes(attribute.String(m.Key(), m.Value()))
	//}
}
