package xray

import (
	"context"
	"os"
	"path/filepath"
)

type ConfigWriter interface {
	Save(ctx context.Context, data []byte) error
}

type FileWriter struct {
	path string
}

func NewConfigWriter(path string) ConfigWriter {
	return &FileWriter{
		path: path,
	}
}

func (w *FileWriter) Save(
	ctx context.Context,
	data []byte,
) error {

	if err := ctx.Err(); err != nil {
		return err
	}

	dir := filepath.Dir(w.path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(
		w.path,
		data,
		0644,
	)
}
