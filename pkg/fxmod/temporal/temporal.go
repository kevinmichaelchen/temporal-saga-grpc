package temporal

import (
	"context"
	"fmt"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/temporal/ctxpropagation"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/opentelemetry"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/fx"
)

var Module = fx.Module("temporal",
	fx.Provide(
		NewClient,
	),
)

func NewClient(lc fx.Lifecycle) (client.Client, error) {
	interceptors, err := clientInterceptors()
	if err != nil {
		return nil, err
	}

	c, err := client.Dial(client.Options{
		Interceptors:       interceptors,
		ContextPropagators: []workflow.ContextPropagator{ctxpropagation.NewContextPropagator()},
	})
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			c.Close()
			return nil
		},
	})

	return c, nil
}

func clientInterceptors() ([]interceptor.ClientInterceptor, error) {
	i, err := opentelemetry.NewTracingInterceptor(opentelemetry.TracerOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create OTEL tracing interceptor: %w", err)
	}
	return []interceptor.ClientInterceptor{
		i,
	}, nil
}
