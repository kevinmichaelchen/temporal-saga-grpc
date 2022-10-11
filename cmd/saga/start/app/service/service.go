package service

import (
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/service"
	"go.temporal.io/sdk/client"
	"go.uber.org/fx"
)

var Module = fx.Module("service",
	fx.Provide(
		NewService,
	),
)

func NewService(c client.Client) *service.Service {
	return service.NewService(c)
}
