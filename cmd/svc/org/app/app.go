// Package app provides an FX module for the application.
package app

import (
	"buf.build/gen/go/kevinmichaelchen/orgapis/connectrpc/go/org/v1beta1/orgv1beta1connect"
	"connectrpc.com/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/sql"
	"go.uber.org/fx"

	modService "github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/service"
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
			path, handler := orgv1beta1connect.NewOrgServiceHandler(
				svc,
				connect.WithInterceptors(pkgConnect.Interceptors()...),
			)

			return modConnect.HandlerOutput{
				Path:    path,
				Handler: handler,
			}
		},
		Services: []string{
			"orgv1beta1.OrgService",
		},
	}),
	logging.Module,
	modService.Module,
	otel.CreateModule(otel.ModuleOptions{
		ServiceName: "org-svc",
	}),
	sql.Module,
)
