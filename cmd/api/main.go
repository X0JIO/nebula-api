package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/X0JIO/nebula-api/internal/app"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	application, err := app.New()
	if err != nil {
		panic(err)
	}

	defer application.Logger.Sync()

	if err := application.Run(ctx); err != nil {
		application.Logger.Fatal(err.Error())
	}
}