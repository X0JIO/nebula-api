package generator

import (
	"fmt"
	"net/url"
)

func GenerateVLESS(identity *Identity, host string) string {
	return fmt.Sprintf(
		"vless://%s@%s:443?encryption=none&security=reality&type=tcp&flow=xtls-rprx-vision&pbk=%s&sid=%s&sni=%s&fp=chrome#Nebula",
		identity.UserUUID,
		host,
		url.QueryEscape(identity.RealityPublicKey),
		identity.RealityShortID,
		host,
	)
}
