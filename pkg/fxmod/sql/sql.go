package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/fx"
	"os"
)

var Module = fx.Module("logging",
	fx.Provide(
		NewConnection,
	),
)

func NewConnection() (*sql.DB, error) {
	conn, err := sql.Open("pgx", os.Getenv("DB_CONN_STRING"))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to db: %w", err)
	}

	return conn, nil
}
