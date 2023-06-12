package app

import (
	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/worker/app/worker"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/temporal"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/tracing"
)

var Module = fx.Options(
	temporal.Module,
	logging.Module,
	worker.Module,
	tracing.CreateModule(tracing.ModuleOptions{
		ServiceName: "temporal-worker",
	}),
)
