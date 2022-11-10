package app

import (
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/app/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/temporal"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/tracing"
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
