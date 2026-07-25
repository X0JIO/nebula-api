package xray

import (
	"os"
	"path/filepath"
)

type ConfigWriter struct {
	path string
}

func NewConfigWriter(
	path string,
) *ConfigWriter {

	return &ConfigWriter{
		path: path,
	}
}

func (w *ConfigWriter) Save(
	cfg []byte,
) error {

	dir := filepath.Dir(w.path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(
		w.path,
		cfg,
		0644,
	)
}
