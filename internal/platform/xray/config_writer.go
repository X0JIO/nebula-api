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

	if err := os.MkdirAll(
		dir,
		0755,
	); err != nil {
		return err
	}

	tmp := w.path + ".tmp"

	file, err := os.Create(tmp)

	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return os.Rename(
		tmp,
		w.path,
	)
}
