package app

import (
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/worker/app/temporal"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/worker/app/tracing"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/worker/app/worker"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"go.uber.org/fx"
)

var Module = fx.Options(
	temporal.Module,
	logging.Module,
	worker.Module,
	tracing.Module,
)
