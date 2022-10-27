package app

import (
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/tracing"
	"github.com/kevinmichaelchen/temporal-saga-grpc/services/temporalstarter/app/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/services/temporalstarter/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/services/temporalstarter/app/temporal"
	"go.uber.org/fx"
)

var Module = fx.Options(
	temporal.Module,
	connect.Module,
	logging.Module,
	service.Module,
	tracing.CreateModule(tracing.ModuleOptions{
		ServiceName: "temporal-starter",
	}),
)
