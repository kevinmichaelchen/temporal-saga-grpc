// Package interceptor provides an FX module for Connect interceptors.
package interceptor

import (
	"connectrpc.com/connect"
	"go.uber.org/fx"

	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
)

// Module - An FX module for Connect interceptors.
var Module = fx.Module("connect/interceptor",
	fx.Provide(
		NewInterceptors(),
	),
)

// NewInterceptors - Creates interceptors for Connect Go servers.
func NewInterceptors() ([]connect.Interceptor, error) {
	return pkgConnect.Interceptors()
}
