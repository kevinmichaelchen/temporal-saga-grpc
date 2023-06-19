// Package app provides an FX module for the application.
package app

import (
	"github.com/bufbuild/connect-go"
	"go.uber.org/fx"

	modService "github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/license/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/license/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/license/v1beta1/licensev1beta1connect"
	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
	modConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/tracing"
)

// Module - An FX module for the application.
var Module = fx.Options(
	modConnect.CreateModule(&modConnect.ModuleOptions{
		HandlerProvider: func(svc *service.Service) modConnect.HandlerOutput {
			// Register our Connect-Go server
			path, handler := licensev1beta1connect.NewLicenseServiceHandler(
				svc,
				connect.WithInterceptors(pkgConnect.UnaryInterceptors()...),
			)

			return modConnect.HandlerOutput{
				Path:    path,
				Handler: handler,
			}
		},
		Services: []string{
			"licensev1beta1.LicenseService",
		},
	}),
	logging.Module,
	modService.Module,
	tracing.CreateModule(tracing.ModuleOptions{
		ServiceName: "license-svc",
	}),
)
