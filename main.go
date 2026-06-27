package main

import (
	"context"
	"finally/features/postgres"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := postgres.Connecting(ctx)
	if err != nil {
		panic(err)
	}

	if err := postgres.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	err = postgres.AddBook(ctx, conn, postgres.Book{
		Title:     "qwewq",
		Author:    "asdasd",
		Review:    nil,
		Date:      2023,
		Date_add:  time.Now(),
		Date_read: nil,
		Is_read:   false,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("succed!")
}
