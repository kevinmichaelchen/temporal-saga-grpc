// Package service provides an FX module for this service's API handlers.
package service

import (
	"github.com/bufbuild/protovalidate-go"
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
func NewService(
	c client.Client,
	validator *protovalidate.Validator,
) *service.Service {
	return service.NewService(c, validator)
}
