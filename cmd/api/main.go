package main

import (
	"context"
	"log"

	"github.com/your-org/nebula-api/internal/bootstrap"
)

func main() {
	ctx := context.Background()

	app, err := bootstrap.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}