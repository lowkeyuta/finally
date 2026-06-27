package postgres

import (
	"time"
)

type Book struct {
	Id            int
	Title, Author string
	Review        *string
	Date          int
	Date_add      time.Time
	Date_read     *time.Time
	Is_read       bool
}
