package generator

import (
	"fmt"
	"net/url"
)

func GenerateVLESS(
	identity *Identity,
	server ServerEndpoint,
) string {
	return fmt.Sprintf(
		"vless://%s@%s:%d?encryption=none&security=reality&type=tcp&flow=xtls-rprx-vision&pbk=%s&sid=%s&sni=%s&fp=chrome#Nebula",
		identity.UserUUID,
		server.Host,
		server.Port,
		url.QueryEscape(server.PublicKey),
		server.ShortID,
		server.Host,
	)
}
