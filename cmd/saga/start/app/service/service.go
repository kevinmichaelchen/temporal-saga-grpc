package service

import (
	"go.temporal.io/sdk/client"
	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/service"
)

// Module - An FX module for the service layer.
var Module = fx.Module("service",
	fx.Provide(
		NewService,
	),
)

// NewService - Constructs a new controller for the service layer.
func NewService(c client.Client) *service.Service {
	return service.NewService(c)
}
