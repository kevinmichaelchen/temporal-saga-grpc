package app

import (
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/rand"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/tracing"
	"github.com/kevinmichaelchen/temporal-saga-grpc/services/license/app/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/services/license/app/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	connect.Module,
	logging.Module,
	service.Module,
	rand.Module,
	tracing.CreateModule(tracing.ModuleOptions{
		ServiceName: "license-svc",
	}),
)
