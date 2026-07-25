package generator

import "fmt"

func GenerateShadowsocks(identity *Identity, host string) string {
	return fmt.Sprintf(
		"ss://%s:%s@%s:443#Nebula-SS",
		"chacha20-ietf-poly1305",
		identity.ShadowsocksPassword,
		host,
	)
}
