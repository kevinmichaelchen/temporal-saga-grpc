// Package app provides an FX module for the application.
package app

import (
	"github.com/bufbuild/connect-go"
	"go.uber.org/fx"

	modService "github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/temporal/v1beta1/temporalv1beta1connect"
	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
	modConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/temporal"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/tracing"
)

// Module - An FX module for the application.
var Module = fx.Options(
	temporal.Module,
	modConnect.CreateModule(&modConnect.ModuleOptions{
		HandlerProvider: func(svc *service.Service) modConnect.HandlerOutput {
			// Register our Connect-Go server
			path, handler := temporalv1beta1connect.NewTemporalServiceHandler(
				svc,
				connect.WithInterceptors(pkgConnect.UnaryInterceptors()...),
			)

			return modConnect.HandlerOutput{
				Path:    path,
				Handler: handler,
			}
		},
		Services: []string{
			"temporalv1beta1.TemporalService",
		},
	}),
	logging.Module,
	modService.Module,
	tracing.CreateModule(tracing.ModuleOptions{
		ServiceName: "temporal-starter",
	}),
)
