package provision

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/X0JIO/nebula-api/internal/platform/xray/config"
)

type DBClientProvider struct {
	queries *db.Queries
}

func NewDBClientProvider(
	queries *db.Queries,
) *DBClientProvider {

	return &DBClientProvider{
		queries: queries,
	}
}

func (p *DBClientProvider) ListClients(
	ctx context.Context,
) ([]config.Client, error) {

	users, err := p.queries.ListVPNUsers(ctx)

	if err != nil {
		return nil, err
	}

	clients := make(
		[]config.Client,
		0,
		len(users),
	)

	for _, user := range users {

		clients = append(
			clients,
			config.Client{
				ID:    user.Uuid,
				Email: user.SubscriptionToken,
			},
		)
	}

	return clients, nil
}
