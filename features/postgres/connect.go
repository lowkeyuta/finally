package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connecting(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:1121122@localhost:5432/postgres")
}
