package xray

import "context"

type Runtime struct {
	process ProcessManager

	config *ConfigService
}

func NewRuntime(
	process ProcessManager,
	config *ConfigService,

) *Runtime {

	return &Runtime{

		process: process,

		config: config,
	}
}

func (r *Runtime) Start(
	ctx context.Context,
) error {

	err := r.config.Generate(ctx)

	if err != nil {

		return err
	}

	return r.process.Start(ctx)
}

func (r *Runtime) Stop(
	ctx context.Context,
) error {

	return r.process.Stop(ctx)
}

func (r *Runtime) Restart(
	ctx context.Context,
) error {

	return r.process.Restart(ctx)
}

func (r *Runtime) Status() Status {

	return r.process.Status()
}
