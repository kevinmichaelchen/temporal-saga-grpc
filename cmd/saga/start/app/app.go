// Package app provides an FX module for the application.
package app

import (
	temporalConnect "buf.build/gen/go/kevinmichaelchen/temporalapis/connectrpc/go/temporal/v1beta1/temporalv1beta1connect"
	"connectrpc.com/connect"
	"go.uber.org/fx"

	modService "github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/service"
	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
	modConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/otel"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/temporal"
)

// Module - An FX module for the application.
var Module = fx.Options(
	temporal.Module,
	modConnect.CreateModule(&modConnect.ModuleOptions{
		HandlerProvider: func(svc *service.Service) modConnect.HandlerOutput {
			// Register our Connect-Go server
			path, handler := temporalConnect.NewTemporalServiceHandler(
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
	otel.CreateModule(otel.ModuleOptions{
		ServiceName: "temporal-starter",
	}),
)
