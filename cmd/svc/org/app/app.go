package app

import (
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/app/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/app/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/app/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	connect.Module,
	logging.Module,
	service.Module,
)
