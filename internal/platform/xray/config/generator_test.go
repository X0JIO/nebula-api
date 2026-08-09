package config

import (
	"encoding/json"
	"testing"

	"github.com/X0JIO/nebula-api/internal/platform/xray/reality"
)

func TestGenerateConfig(t *testing.T) {

	credentials, err := reality.Generate()

	if err != nil {
		t.Fatal(err)
	}

	generator := NewGenerator(credentials)

	client := generator.GenerateClient()

	data, err := generator.GenerateConfig(client)

	if err != nil {
		t.Fatal(err)
	}

	var result map[string]any

	dataBytes := data

	if err := json.Unmarshal(
		dataBytes,
		&result,
	); err != nil {
		t.Fatal(err)
	}

	if result["inbounds"] == nil {
		t.Fatal("missing inbounds")
	}

}
