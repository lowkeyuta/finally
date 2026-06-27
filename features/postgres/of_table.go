package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func AddBook(ctx context.Context, conn *pgx.Conn, book Book) error {
	sqlQuery := `
	INSERT INTO books(title,author,review,date,date_add,date_read,is_read)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := conn.Exec(ctx, sqlQuery,
		book.Title,
		book.Author,
		book.Review,
		book.Date,
		book.Date_add,
		book.Date_read,
		book.Is_read)

	return err
}
