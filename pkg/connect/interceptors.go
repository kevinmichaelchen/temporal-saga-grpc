// Package connect returns Connect Go interceptors.
package connect

import (
	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
)

// Interceptors - Interceptors for Connect Go servers.
func Interceptors() []connect.Interceptor {
	return []connect.Interceptor{
		otelconnect.NewInterceptor(),
	}
}
