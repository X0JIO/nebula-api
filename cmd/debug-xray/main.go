package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/X0JIO/nebula-api/internal/platform/xray/config"
)

func main() {
	cfg := config.DefaultVPNConfig()

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	if err := os.WriteFile("xray.generated.json", data, 0644); err != nil {
		panic(err)
	}

	fmt.Println("generated xray.generated.json")
}
