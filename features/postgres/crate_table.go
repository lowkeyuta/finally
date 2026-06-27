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
	date INT NOT NULL,
	date_read TIMESTAMP ,
	date_add TIMESTAMP NOT NULL,
	is_read BOOLEAN);
	`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
