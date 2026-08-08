package generator

import "fmt"

func GenerateShadowsocks(
	identity *Identity,
	server ServerEndpoint,
) string {
	return fmt.Sprintf(
		"ss://%s:%s@%s:%d#Nebula-SS",
		"chacha20-ietf-poly1305",
		identity.ShadowsocksPassword,
		server.Host,
		server.Port,
	)
}
