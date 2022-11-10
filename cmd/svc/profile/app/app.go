package app

import (
	"github.com/bufbuild/connect-go"
	modService "github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/app/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/com/teachingstrategies/profile/v1beta1/profilev1beta1connect"
	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
	modConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/logging"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/rand"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/tracing"
	"go.uber.org/fx"
)

var Module = fx.Options(
	modConnect.CreateModule(&modConnect.ModuleOptions{
		HandlerProvider: func(svc *service.Service) modConnect.HandlerOutput {
			// Register our Connect-Go server
			path, h := profilev1beta1connect.NewProfileServiceHandler(
				svc,
				connect.WithInterceptors(pkgConnect.UnaryInterceptors()...),
			)
			return modConnect.HandlerOutput{
				Path:    path,
				Handler: h,
			}
		},
		Services: []string{
			"com.teachingstrategies.profilev1beta1.ProfileService",
		},
	}),
	logging.Module,
	modService.Module,
	rand.Module,
	tracing.CreateModule(tracing.ModuleOptions{
		ServiceName: "profile-svc",
	}),
)
