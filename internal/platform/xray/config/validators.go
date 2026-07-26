package config

import (
	"fmt"

	"github.com/X0JIO/nebula-api/internal/platform/xray/model"
)

func ValidateInbound(i model.Inbound) error {
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
