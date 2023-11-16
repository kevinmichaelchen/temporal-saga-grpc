// Package app provides an FX module for the application.
package app

import (
	licenseConnect "buf.build/gen/go/kevinmichaelchen/licenseapis/connectrpc/go/license/v1beta1/licensev1beta1connect"
	"connectrpc.com/connect"
	"go.uber.org/fx"

	modService "github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/license/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/license/service"
	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
	modConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/otel"
)

// Module - An FX module for the application.
var Module = fx.Options(
	modConnect.CreateModule(&modConnect.ModuleOptions{
		HandlerProvider: func(svc *service.Service) modConnect.HandlerOutput {
			// Register our Connect-Go server
			path, handler := licenseConnect.NewLicenseServiceHandler(
				svc,
				connect.WithInterceptors(pkgConnect.Interceptors()...),
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
	otel.CreateModule(otel.ModuleOptions{
		ServiceName: "license-svc",
	}),
)
