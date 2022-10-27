package app

import (
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/tracing"
	"github.com/kevinmichaelchen/temporal-saga-grpc/services/temporalworker/app/temporal"
	"github.com/kevinmichaelchen/temporal-saga-grpc/services/temporalworker/app/worker"
	"go.uber.org/fx"
)

var Module = fx.Options(
	temporal.Module,
	logging.Module,
	worker.Module,
	tracing.CreateModule(tracing.ModuleOptions{
		ServiceName: "temporal-worker",
	}),
)
