// Package app provides an FX module for the application.
package app

import (
	profilev1beta1connect "buf.build/gen/go/kevinmichaelchen/profileapis/connectrpc/go/profile/v1beta1/profilev1beta1connect"
	"connectrpc.com/connect"
	"go.uber.org/fx"

	modService "github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/service"
	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
	modConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/otel"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/sql"
)

// Module - An FX module for the application.
var Module = fx.Options(
	modConnect.CreateModule(&modConnect.ModuleOptions{
		HandlerProvider: func(svc *service.Service) modConnect.HandlerOutput {
			// Register our Connect-Go server
			path, handler := profilev1beta1connect.NewProfileServiceHandler(
				svc,
				connect.WithInterceptors(pkgConnect.Interceptors()...),
			)

			return modConnect.HandlerOutput{
				Path:    path,
				Handler: handler,
			}
		},
		Services: []string{
			"profilev1beta1.ProfileService",
		},
	}),
	logging.Module,
	modService.Module,
	otel.CreateModule(otel.ModuleOptions{
		ServiceName: "profile-svc",
	}),
	sql.Module,
)
