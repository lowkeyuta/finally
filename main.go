package main

import (
	"context"
	"finally/features/postgres"
	"fmt"
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

	fmt.Println("succed!")
}
