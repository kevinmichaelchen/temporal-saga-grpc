// Package temporal provides interceptors and middleware for the Temporal
// client.
package temporal

import (
	"fmt"

	"go.temporal.io/sdk/contrib/opentelemetry"
	"go.temporal.io/sdk/interceptor"
)

// ClientInterceptors - Interceptors for Temporal clients.
func ClientInterceptors() ([]interceptor.ClientInterceptor, error) {
	tracingInterceptor, err := opentelemetry.NewTracingInterceptor(
		opentelemetry.TracerOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create OTEL tracing interceptor: %w", err)
	}

	return []interceptor.ClientInterceptor{
		tracingInterceptor,
	}, nil
}
