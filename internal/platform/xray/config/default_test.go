package config_test

import (
	"testing"

	"github.com/X0JIO/nebula-api/internal/platform/xray/config"
)

func TestDefaultVPNConfig(t *testing.T) {

	cfg := config.DefaultVPNConfig()

	if len(cfg.Inbounds) == 0 {
		t.Fatal("expected inbounds")
	}

	if len(cfg.Outbounds) == 0 {
		t.Fatal("expected outbounds")
	}

	t.Logf(
		"inbounds=%d outbounds=%d",
		len(cfg.Inbounds),
		len(cfg.Outbounds),
	)
}
