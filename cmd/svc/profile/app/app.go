package app

import (
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/app/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/app/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/app/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	connect.Module,
	logging.Module,
	service.Module,
)
