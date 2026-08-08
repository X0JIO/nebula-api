package generator

import (
	"fmt"
	"net/url"
)

func GenerateReality(
	identity *Identity,
	server ServerEndpoint,
) string {
	return fmt.Sprintf(
		"vless://%s@%s:%d?security=reality&encryption=none&type=tcp&flow=xtls-rprx-vision&pbk=%s&sid=%s&sni=%s&fp=chrome#Nebula-Reality",
		identity.UserUUID,
		server.Host,
		server.Port,
		url.QueryEscape(server.PublicKey),
		server.ShortID,
		server.Host,
	)
}
