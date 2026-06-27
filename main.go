package main

import (
	"bufio"
	"context"
	"finally/features/postgres"
	"fmt"
	"os"
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

	if Verif() {
		err := postgres.AddBook(ctx, conn, postgres.Book{
			Title:     "",
			Author:    "",
			Review:    nil,
			Date:      3333,
			Date_add:  time.Now(),
			Date_read: nil,
			Is_read:   false,
		})
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("succed!")
}

func Verif() bool {
	ch := make(chan struct{})
	fmt.Println("yes or no?")
	go func() {
		bufio.NewReader(os.Stdin).ReadString('\n')
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		return true
	case <-time.After(3 * time.Second):
		return false
	}
}
