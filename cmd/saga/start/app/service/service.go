package service

import (
	"go.temporal.io/sdk/client"
	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/service"
)

var Module = fx.Module("service",
	fx.Provide(
		NewService,
	),
)

func NewService(c client.Client) *service.Service {
	return service.NewService(c)
}
