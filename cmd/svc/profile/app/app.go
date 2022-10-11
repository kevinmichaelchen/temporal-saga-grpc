package app

import (
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/app/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/rand"
	"go.uber.org/fx"
)

var Module = fx.Options(
	connect.Module,
	logging.Module,
	service.Module,
	rand.Module,
)
