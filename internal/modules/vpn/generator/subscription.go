package generator

import (
	"encoding/base64"
	"strings"
)

func GenerateSubscription(
	vless string,
	reality string,
	vmess string,
	trojan string,
	shadowsocks string,
) string {

	configs := []string{
		vless,
		reality,
		vmess,
		trojan,
		shadowsocks,
	}

	lines := make([]string, 0, len(configs))

	for _, cfg := range configs {
		if cfg != "" {
			lines = append(lines, cfg)
		}
	}

	return strings.Join(lines, "\n")
}

func GenerateSubscriptionBase64(
	vless string,
	reality string,
	vmess string,
	trojan string,
	shadowsocks string,
) string {

	text := GenerateSubscription(
		vless,
		reality,
		vmess,
		trojan,
		shadowsocks,
	)

	return base64.StdEncoding.EncodeToString([]byte(text))
}
