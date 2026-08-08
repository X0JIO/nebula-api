package generator

import "fmt"

func GenerateTrojan(
	identity *Identity,
	server ServerEndpoint,
) string {
	return fmt.Sprintf(
		"trojan://%s@%s:%d?security=tls&sni=%s&type=tcp#Nebula-Trojan",
		identity.UserUUID,
		server.Host,
		server.Port,
		server.Host,
	)
}
