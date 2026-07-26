package xray

import "context"

type Service struct {
	process ProcessManager
	config  *ConfigService
}

func NewService(
	process ProcessManager,
	config *ConfigService,
) *Service {

	return &Service{
		process: process,
		config:  config,
	}
}

func (s *Service) Start(
	ctx context.Context,
) error {

	if s.config != nil {

		if err := s.config.Generate(ctx); err != nil {
			return err
		}
	}

	return s.process.Start(ctx)
}

func (s *Service) Stop(
	ctx context.Context,
) error {

	return s.process.Stop(ctx)
}

func (s *Service) Restart(
	ctx context.Context,
) error {

	if err := s.Stop(ctx); err != nil {
		return err
	}

	return s.Start(ctx)
}

func (s *Service) Running() bool {

	return s.process.Running()
}

func (s *Service) PID() int {

	return s.process.PID()
}

func (s *Service) Status() Status {

	return Status{
		Running: s.Running(),
		PID:     s.PID(),
	}
}
