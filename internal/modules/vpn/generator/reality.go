package generator

import (
	"fmt"
	"net/url"
)

func GenerateReality(identity *Identity, host string) string {
	return fmt.Sprintf(
		"vless://%s@%s:443?security=reality&encryption=none&type=tcp&flow=xtls-rprx-vision&pbk=%s&sid=%s&sni=%s&fp=chrome#Nebula-Reality",
		identity.UserUUID,
		host,
		url.QueryEscape(identity.RealityPublicKey),
		identity.RealityShortID,
		host,
	)
}
