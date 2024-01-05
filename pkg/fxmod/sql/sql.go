// Package sql provides an FX module for SQL connections.
package sql

import (
	"database/sql"
	"fmt"
	"os"

	"go.uber.org/fx"

	// Blank import is okay here.
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Module - An FX module for a SQL connection.
var Module = fx.Module("logging",
	fx.Provide(
		NewConnection,
	),
)

// NewConnection - Returns a new SQL connection.
func NewConnection() (*sql.DB, error) {
	conn, err := sql.Open("pgx", os.Getenv("DB_CONN_STRING"))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to db: %w", err)
	}

	return conn, nil
}
