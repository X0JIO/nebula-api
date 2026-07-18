package main

import (
	"os"

	"github.com/X0JIO/nebula-api/internal/app"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	application, err := app.New()
	if err != nil {
		panic(err)
	}

	defer application.Logger.Sync()

	application.Logger.Info("application starting")

	application.Logger.Info(
		"configuration loaded",
	)

	application.Logger.Info(
		"environment",
	)

	application.Logger.Info(
		application.Config.App.Env,
	)

	application.Logger.Info("application initialized")

	os.Exit(0)

}