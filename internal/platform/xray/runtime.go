package xray

import (
	"context"
	"os"
	"path/filepath"
)

type Runtime struct {
	ConfigService *ConfigService

	Process *Service

	BinaryPath string

	ConfigPath string
}

func NewRuntime(
	baseDir string,
	binaryPath string,
) *Runtime {

	configPath := filepath.Join(
		baseDir,
		"config.json",
	)

	writer := NewConfigWriter(
		configPath,
	)

	configService := NewConfigService(
		writer,
	)

	process := NewProcessManager(
		binaryPath,
		configPath,
		baseDir,
	)

	return &Runtime{

		ConfigService: configService,

		Process: NewService(process),

		BinaryPath: binaryPath,

		ConfigPath: configPath,
	}
}

func (r *Runtime) Prepare() error {

	dir := filepath.Dir(
		r.ConfigPath,
	)

	return os.MkdirAll(
		dir,
		0755,
	)
}

func (r *Runtime) Start(
	ctx context.Context,
) error {

	if err := r.Prepare(); err != nil {
		return err
	}

	return r.Process.Start(ctx)
}
