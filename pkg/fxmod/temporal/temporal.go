package temporal

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/client"
	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/temporal"
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

	temporalClient, err := client.Dial(client.Options{
		Interceptors: interceptors,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to dial Temporal client: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: nil,
		OnStop: func(ctx context.Context) error {
			temporalClient.Close()

			return nil
		},
	})

	return temporalClient, nil
}
