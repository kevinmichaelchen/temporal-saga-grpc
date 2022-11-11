package temporal

import (
	"context"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/temporal"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/temporal/ctxpropagation"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/fx"
)

var Module = fx.Module("temporal",
	fx.Provide(
		NewClient,
	),
)

func NewClient(lc fx.Lifecycle) (client.Client, error) {
	interceptors, err := temporal.ClientInterceptors()
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
