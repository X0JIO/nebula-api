package config

import (
	"encoding/json"

	"github.com/X0JIO/nebula-api/internal/platform/xray/reality"
)

type Generator struct {
	Reality *reality.Credentials
}

func (g *Generator) GenerateConfig(
	client Client,
) ([]byte, error) {

	return g.Generate(
		[]Client{
			client,
		},
	)
}

func NewGenerator(
	credentials *reality.Credentials,
) *Generator {

	return &Generator{
		Reality: credentials,
	}
}

func (g *Generator) GenerateClient() Client {

	return Client{
		ID:    g.Reality.UUID,
		Email: g.Reality.UUID + "@nebula",
	}
}

func (g *Generator) Generate(
	clients []Client,
) ([]byte, error) {

	xrayClients := make([]any, 0, len(clients))

	for _, client := range clients {

		xrayClients = append(
			xrayClients,
			map[string]any{
				"id":    client.ID,
				"email": client.Email,
			},
		)
	}

	cfg := map[string]any{

		"log": map[string]any{
			"loglevel": "warning",
		},

		"inbounds": []any{

			map[string]any{

				"port": 443,

				"protocol": "vless",

				"settings": map[string]any{

					"clients": xrayClients,

					"decryption": "none",
				},

				"streamSettings": map[string]any{

					"network": "tcp",

					"security": "reality",

					"realitySettings": map[string]any{

						"show": false,

						"dest": "www.cloudflare.com:443",

						"serverNames": []string{
							"www.cloudflare.com",
						},

						"privateKey": g.Reality.PrivateKey,

						"shortIds": []string{
							g.Reality.ShortID,
						},
					},
				},
			},
		},

		"outbounds": []any{

			map[string]any{
				"protocol": "freedom",
			},
		},
	}

	return json.MarshalIndent(
		cfg,
		"",
		"  ",
	)
}
