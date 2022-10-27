package temporal

import (
	"context"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/temporal"
	"go.temporal.io/sdk/client"
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
		Interceptors: interceptors,
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
