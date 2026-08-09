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

// Deploy generates and writes Xray configuration.
func (s *Service) Deploy(
	ctx context.Context,
) error {

	if s.config == nil {
		return nil
	}

	return s.config.Generate(ctx)
}

// DeployAndRestart applies new config and restarts Xray.
func (s *Service) DeployAndRestart(
	ctx context.Context,
) error {

	if err := s.Deploy(ctx); err != nil {
		return err
	}

	return s.Restart(ctx)
}

func (s *Service) Start(
	ctx context.Context,
) error {

	if err := s.Validate(ctx); err != nil {
		return err
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

	if s.client != nil {

		if err := s.client.Reload(ctx); err == nil {
			return nil
		}
	}

	return s.process.Restart(ctx)
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
