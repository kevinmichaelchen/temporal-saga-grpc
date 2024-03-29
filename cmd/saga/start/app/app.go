// Package app provides an FX module for the application.
package app

import (
	temporalConnect "buf.build/gen/go/kevinmichaelchen/temporalapis/connectrpc/go/temporal/v1beta1/temporalv1beta1connect"
	"connectrpc.com/connect"
	"go.uber.org/fx"

	modService "github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/service"
	modConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/otel"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/temporal"
)

// Module - An FX module for the application.
var Module = fx.Options(
	temporal.Module,
	modConnect.CreateModule(&modConnect.ModuleOptions{
		HandlerProvider: func(svc *service.Service, interceptors []connect.Interceptor) modConnect.HandlerOutput {
			// Register our Connect-Go server
			path, handler := temporalConnect.NewTemporalServiceHandler(
				svc,
				connect.WithInterceptors(interceptors...),
			)

			return modConnect.HandlerOutput{
				Path:    path,
				Handler: handler,
			}
		},
		Service: temporalConnect.TemporalServiceName,
	}),
	logging.Module,
	modService.Module,
	otel.CreateModule(otel.ModuleOptions{
		ServiceName: "temporal-starter",
	}),
)
