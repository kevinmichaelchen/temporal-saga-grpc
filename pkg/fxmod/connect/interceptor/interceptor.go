package interceptor

import (
	"connectrpc.com/connect"
	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
	"go.uber.org/fx"
)

var Module = fx.Module("connect/interceptor",
	fx.Provide(
		NewInterceptors(),
	),
)

// NewInterceptors - Creates interceptors for Connect Go servers.
func NewInterceptors() ([]connect.Interceptor, error) {
	return pkgConnect.Interceptors()
}
