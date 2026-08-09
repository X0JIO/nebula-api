package config

import "context"

type ClientProvider interface {
	ListClients(
		ctx context.Context,
	) ([]Client, error)
}

type Deployer struct {
	generator *Generator

	writer interface {
		Save(
			ctx context.Context,
			data []byte,
		) error
	}

	clients ClientProvider
}

func NewDeployer(
	generator *Generator,
	writer interface {
		Save(context.Context, []byte) error
	},
	clients ClientProvider,
) *Deployer {

	return &Deployer{
		generator: generator,
		writer:    writer,
		clients:   clients,
	}
}

func (d *Deployer) Deploy(
	ctx context.Context,
) error {

	clients, err := d.clients.ListClients(ctx)

	if err != nil {
		return err
	}

	data, err := d.generator.Generate(
		clients,
	)

	if err != nil {
		return err
	}

	return d.writer.Save(
		ctx,
		data,
	)
}
