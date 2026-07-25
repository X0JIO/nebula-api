package config

import "fmt"

func ValidateInbound(i Inbound) error {
	if i.Tag == "" {
		return fmt.Errorf("tag is empty")
	}

	if i.Port <= 0 {
		return fmt.Errorf("invalid port")
	}

	if i.Protocol == "" {
		return fmt.Errorf("protocol is empty")
	}

	return nil
}
