package config

import (
	"encoding/json"

	"github.com/X0JIO/nebula-api/internal/platform/xray/reality"
)

type Generator struct {
	Reality *reality.Credentials
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

func (g *Generator) GenerateConfig(
	client Client,
) ([]byte, error) {

	cfg := map[string]any{

		"log": map[string]any{
			"loglevel": "warning",
		},

		"inbounds": []any{
			map[string]any{

				"port": 443,

				"protocol": "vless",

				"settings": map[string]any{

					"clients": []any{
						map[string]any{

							"id": client.ID,

							"email": client.ID,
						},
					},

					"decryption": "none",
				},

				"streamSettings": map[string]any{

					"network": "tcp",

					"security": "reality",
					"realitySettings": map[string]any{

						"show": false,

						"dest": "www.cloudflare.com:443",

						"xver": 0,

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

		"nebula": map[string]any{

			"publicKey": g.Reality.PublicKey,

			"shortId": g.Reality.ShortID,
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
