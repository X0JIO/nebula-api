package config

import "encoding/json"

func Generate(cfg Config) ([]byte, error) {
	return json.MarshalIndent(
		cfg,
		"",
		"  ",
	)
}

func MustGenerate(cfg Config) []byte {
	data, err := Generate(cfg)
	if err != nil {
		panic(err)
	}
	return data
}
