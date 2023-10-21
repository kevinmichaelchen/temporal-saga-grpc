// Package app provides an FX module for the application.
package app

import (
	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/worker/app/worker"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/otel"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/temporal"
)

// Module - An FX module for the application.
var Module = fx.Options(
	temporal.Module,
	logging.Module,
	worker.Module,
	otel.CreateModule(otel.ModuleOptions{
		ServiceName: "temporal-worker",
	}),
)
