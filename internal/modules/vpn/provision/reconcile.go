package provision

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/X0JIO/nebula-api/internal/platform/xray"
	"github.com/X0JIO/nebula-api/internal/platform/xray/config"
)

type Reconciler struct {
	queries   *db.Queries
	generator *config.Generator
	writer    xray.ConfigWriter
	xray      xray.Controller
}

func NewReconciler(
	queries *db.Queries,
	generator *config.Generator,
	writer xray.ConfigWriter,
	controller xray.Controller,
) *Reconciler {

	return &Reconciler{
		queries:   queries,
		generator: generator,
		writer:    writer,
		xray:      controller,
	}
}

func (r *Reconciler) Sync(
	ctx context.Context,
) error {

	users, err := r.queries.ListVPNUsers(ctx)

	if err != nil {
		return err
	}

	clients := make([]config.Client, 0, len(users))

	for _, user := range users {

		clients = append(
			clients,
			config.Client{
				ID:    user.Uuid,
				Email: user.SubscriptionToken,
			},
		)
	}

	cfg, err := r.generator.Generate(
		clients,
	)

	if err != nil {
		return err
	}

	if err := r.writer.Save(
		ctx,
		cfg,
	); err != nil {
		return err
	}

	if err := r.xray.Validate(ctx); err != nil {
		return err
	}

	return r.xray.Reload(ctx)
}
