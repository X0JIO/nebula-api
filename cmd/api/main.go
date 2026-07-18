package main

import (
	"fmt"
	"log"

	"github.com/X0JIO/nebula-api/internal/platform/config"
	"github.com/joho/godotenv"
)

func main() {

	// Загружаем .env если файл существует.
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"%s started (%s)\n",
		cfg.App.Name,
		cfg.App.Env,
	)
}