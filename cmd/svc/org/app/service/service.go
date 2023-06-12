package service

import (
	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/service"
)

var Module = fx.Module("service",
	fx.Provide(
		NewService,
	),
)

func NewService() *service.Service {
	return service.NewService()
}
