package temporal

import (
	"context"
	"go.temporal.io/sdk/client"
	"go.uber.org/fx"
)

var Module = fx.Module("temporal",
	fx.Provide(
		NewClient,
	),
)

func NewClient(lc fx.Lifecycle) (client.Client, error) {
	c, err := client.Dial(client.Options{})
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			c.Close()
			return nil
		},
	})

	return c, nil
}
