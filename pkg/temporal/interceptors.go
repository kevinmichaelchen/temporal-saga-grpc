package temporal

import (
	"fmt"

	"go.temporal.io/sdk/contrib/opentelemetry"
	"go.temporal.io/sdk/interceptor"
)

func ClientInterceptors() ([]interceptor.ClientInterceptor, error) {
	tracingInterceptor, err := opentelemetry.NewTracingInterceptor(
		opentelemetry.TracerOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTEL tracing interceptor: %w", err)
	}

	return []interceptor.ClientInterceptor{
		tracingInterceptor,
	}, nil
}
