package generator

import "fmt"

func GenerateTrojan(identity *Identity, host string) string {
	return fmt.Sprintf(
		"trojan://%s@%s:443?security=tls&sni=%s&type=tcp#Nebula-Trojan",
		identity.UserUUID,
		host,
		host,
	)
}
