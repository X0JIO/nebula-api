package xray

import "context"

type Service struct {
	process *ProcessManager
}

func NewService(
	process *ProcessManager,
) *Service {

	return &Service{
		process: process,
	}
}

func (s *Service) Start(
	ctx context.Context,
) error {

	return s.process.Start(ctx)
}

func (s *Service) Stop() error {

	return s.process.Stop()
}

func (s *Service) Restart(
	ctx context.Context,
) error {

	return s.process.Restart(ctx)
}

func (s *Service) Running() bool {

	return s.process.Running()
}
