// Package main provides the entrypoint for the Profile Service.
package main

import (
	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/app"
)

func main() {
	a := fx.New(
		app.Module,
		// TODO configure logger
		// For details, see https://github.com/uber-go/fx/blob/master/fxevent/zap.go
	)
	a.Run()
}
