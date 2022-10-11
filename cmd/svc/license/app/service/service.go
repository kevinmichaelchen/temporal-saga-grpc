package service

import (
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/license/service"
	"go.uber.org/fx"
)

var Module = fx.Module("service",
	fx.Provide(
		NewService,
	),
)

func NewService() *service.Service {
	return service.NewService()
}
