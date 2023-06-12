package app

import (
	"github.com/bufbuild/connect-go"
	"go.buf.build/bufbuild/connect-go/kevinmichaelchen/licenseapis/license/v1beta1/licensev1beta1connect"
	"go.uber.org/fx"

	modService "github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/license/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/license/service"
	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
	modConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/rand"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/tracing"
)

var Module = fx.Options(
	modConnect.CreateModule(&modConnect.ModuleOptions{
		HandlerProvider: func(svc *service.Service) modConnect.HandlerOutput {
			// Register our Connect-Go server
			path, h := licensev1beta1connect.NewLicenseServiceHandler(
				svc,
				connect.WithInterceptors(pkgConnect.UnaryInterceptors()...),
			)
			return modConnect.HandlerOutput{
				Path:    path,
				Handler: h,
			}
		},
		Services: []string{
			"licensev1beta1.LicenseService",
		},
	}),
	logging.Module,
	modService.Module,
	rand.Module,
	tracing.CreateModule(tracing.ModuleOptions{
		ServiceName: "license-svc",
	}),
)
