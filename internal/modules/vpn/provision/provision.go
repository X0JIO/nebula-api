package provision

import (
	"context"
)

func (s *Service) Provision(
	ctx context.Context,
) error {

	client := s.generator.GenerateClient()

	cfg, err := s.generator.GenerateConfig(
		client,
	)

	if err != nil {
		return err
	}

	if err := s.writer.Save(
		ctx,
		cfg,
	); err != nil {
		return err
	}

	if err := s.xray.Reload(ctx); err != nil {
		return err
	}

	return nil
}
