// Package interceptor provides an FX module for Connect interceptors.
package interceptor

import (
	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"fmt"
	"go.uber.org/fx"
)

// Module - An FX module for Connect interceptors.
var Module = fx.Module("connect/interceptor",
	fx.Provide(
		NewInterceptors,
	),
)

// NewInterceptors - Creates interceptors for Connect Go servers.
func NewInterceptors() ([]connect.Interceptor, error) {
	interceptor, err := otelconnect.NewInterceptor(
		// Trust the client's tracing information. With this option, servers
		// will create child spans for each request.
		otelconnect.WithTrustRemote(),
		// Reduce metrics and tracing cardinality. Drop high-cardinality
		// attributes like the server IP address and the remote port number.
		otelconnect.WithoutServerPeerAttributes(),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to construct otelconnect interceptor: %w", err)
	}

	return []connect.Interceptor{
		interceptor,
	}, nil
}
