// Package main provides the entrypoint for the Org Service.
package main

import (
	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/app"
)

func main() {
	a := fx.New(
		app.Module,
		// TODO configure logger
		// For details, see https://github.com/uber-go/fx/blob/master/fxevent/zap.go
	)
	a.Run()
}
