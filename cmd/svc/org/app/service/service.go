// Package service provides an FX module for this service's API handlers.
package service

import (
	"database/sql"

	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/service"
)

// Module - An FX module for the service layer.
var Module = fx.Module("service",
	fx.Provide(
		NewService,
	),
)

// NewService - Returns a new controller for our business logic.
func NewService(db *sql.DB) *service.Service {
	return service.NewService(db)
}
