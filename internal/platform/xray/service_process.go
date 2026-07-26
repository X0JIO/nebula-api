package xray

import "context"

type Service struct {
	process ProcessManager
	client  Client
	config  *ConfigService
}

func NewService(
	process ProcessManager,
	client Client,
	config *ConfigService,
) *Service {

	return &Service{
		process: process,
		client:  client,
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

	return s.process.Restart(ctx)
}

func (s *Service) Reload(
	ctx context.Context,
) error {

	if s.client == nil {
		return nil
	}

	return s.client.Reload(ctx)
}

func (s *Service) Version(
	ctx context.Context,
) (string, error) {

	if s.client == nil {
		return "unknown", nil
	}

	return s.client.Version(ctx)
}

func (s *Service) Status() Status {

	return s.process.Status()
}

func (s *Service) Validate(
	ctx context.Context,
) error {

	if s.client == nil {
		return nil
	}

	return s.client.Validate(ctx)
}
