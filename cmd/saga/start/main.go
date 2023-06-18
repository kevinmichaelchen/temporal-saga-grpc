// Package main provides the entrypoint for the Temporal Workflow starter API,
// a gRPC/Connect/HTTP API you can use to kick off workflows.
package main

import (
	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/app"
)

func main() {
	a := fx.New(
		app.Module,
		// TODO configure logrus
		// For details, see https://github.com/uber-go/fx/blob/master/fxevent/zap.go
	)
	a.Run()
}
