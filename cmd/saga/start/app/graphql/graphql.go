package graphql

import (
	"context"
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Module("graphql",
	fx.Provide(
		NewSchema,
		NewHandler,
	),
	fx.Invoke(
		Register,
	),
)

func NewSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}

func NewHandler(lc fx.Lifecycle, schema graphql.Schema) *handler.Handler {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			address := ":9094"
			go func() {
				err := http.ListenAndServe(address, nil)
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					logrus.WithError(err).Error("graphql-go ListenAndServe failed")
				}
			}()
			logrus.WithField("address", address).Info("Listening for GraphQL")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return h
}

func Register(h *handler.Handler) {
	http.Handle("/graphql", h)
}
