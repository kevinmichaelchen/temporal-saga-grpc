// Package temporal provides an FX module for a Temporal client.
package temporal

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/client"
	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/temporal"
)

// Module - An FX module for a Temporal client.
var Module = fx.Module("temporal",
	fx.Provide(
		NewClient,
	),
)

// NewClient - Returns a new Temporal client.
//
//nolint:ireturn
func NewClient(lifecycle fx.Lifecycle) (client.Client, error) {
	interceptors, err := temporal.ClientInterceptors()
	if err != nil {
		return nil, err
	}

	temporalClient, err := client.Dial(
		client.Options{
			Interceptors: interceptors,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("unable to dial Temporal client: %w", err)
	}

	lifecycle.Append(fx.Hook{
		OnStart: nil,
		OnStop: func(ctx context.Context) error {
			temporalClient.Close()

			return nil
		},
	})

	return temporalClient, nil
}
