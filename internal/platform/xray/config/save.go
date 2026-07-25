package config

import "os"

func Save(path string, cfg Config) error {

	data, err := Generate(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(
		path,
		data,
		0644,
	)
}
