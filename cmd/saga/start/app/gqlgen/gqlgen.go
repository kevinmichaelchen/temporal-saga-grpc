package gqlgen

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/graph"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/graph/generated"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"net/http"
)

const defaultPort = 9094

var Module = fx.Module("gqlgen",
	fx.Provide(
		NewServer,
	),
	fx.Invoke(
		RegisterServer,
	),
)

func NewServer(lc fx.Lifecycle) *handler.Server {
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			address := fmt.Sprintf(":%d", defaultPort)
			go func() {
				err := http.ListenAndServe(address, nil)
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					logrus.WithError(err).Error("GraphQL (gqlgen) ListenAndServe failed")
				}
			}()
			logrus.WithField("address", address).Info("Listening for GraphQL (gqlgen)")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return srv
}

func RegisterServer(srv *handler.Server) {
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
}
