// Package app provides an FX module for the application.
package app

import (
	"github.com/bufbuild/connect-go"
	"go.buf.build/bufbuild/connect-go/kevinmichaelchen/orgapis/org/v1beta1/orgv1beta1connect"
	"go.uber.org/fx"

	modService "github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/service"
	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
	modConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/rand"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/tracing"
)

// Module - An FX module for the application.
var Module = fx.Options(
	modConnect.CreateModule(&modConnect.ModuleOptions{
		HandlerProvider: func(svc *service.Service) modConnect.HandlerOutput {
			// Register our Connect-Go server
			path, handler := orgv1beta1connect.NewOrgServiceHandler(
				svc,
				connect.WithInterceptors(pkgConnect.UnaryInterceptors()...),
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
	rand.Module,
	tracing.CreateModule(tracing.ModuleOptions{
		ServiceName: "org-svc",
	}),
)
