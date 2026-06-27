package main

import (
	"context"
	"finally/features/postgres"
)

func main() {
	ctx := context.Background()

	_, err := postgres.Connecting(ctx)
	if err != nil {
		panic(err)
	}

}
