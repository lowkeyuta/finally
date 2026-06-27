package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books(
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
	author VARCHAR(50) NOT NULL,
	review VARCHAR(500),
	date TIMESTAMP,
	date_read TIMESTAMP NOT NULL,
	is_read BOOLEAN,
	is_add TIMESTAMP NOT NULL);
	`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
