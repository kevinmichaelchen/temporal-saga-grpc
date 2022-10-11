package app

import (
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/worker/app/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/worker/app/temporal"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/worker/app/worker"
	"go.uber.org/fx"
)

var Module = fx.Options(
	temporal.Module,
	logging.Module,
	worker.Module,
)
