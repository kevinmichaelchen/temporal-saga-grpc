// Package connect returns Connect Go interceptors.
package connect

import (
	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
)

// Interceptors - Interceptors for Connect Go servers.
func Interceptors() []connect.Interceptor {
	return []connect.Interceptor{
		otelconnect.NewInterceptor(
			// Trust the client's tracing information. With this option, servers
			// will create child spans for each request.
			otelconnect.WithTrustRemote(),
			// Reduce metrics and tracing cardinality. Drop high-cardinality
			// attributes like the server IP address and the remote port number.
			otelconnect.WithoutServerPeerAttributes(),
		),
	}
}
