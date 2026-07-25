package xray

import (
	"context"
	"fmt"
)

type InboundManager struct {
	client Client
}

func NewInboundManager(client Client) *InboundManager {
	return &InboundManager{
		client: client,
	}
}

func (m *InboundManager) EnsureInbound(
	ctx context.Context,
	tag string,
	port int,
	protocol string,
) error {

	_, err := m.client.FindInboundByProtocol(
		ctx,
		protocol,
	)

	if err == nil {
		return nil
	}

	return m.client.CreateInbound(
		ctx,
		CreateInboundRequest{
			Tag:      tag,
			Port:     port,
			Protocol: protocol,
		},
	)
}

func (m *InboundManager) EnsureDefaults(
	ctx context.Context,
) error {

	defaults := []CreateInboundRequest{

		{
			Tag:      "vless-reality",
			Port:     443,
			Protocol: "vless",
		},

		{
			Tag:      "vmess",
			Port:     10086,
			Protocol: "vmess",
		},

		{
			Tag:      "trojan",
			Port:     8443,
			Protocol: "trojan",
		},

		{
			Tag:      "shadowsocks",
			Port:     8388,
			Protocol: "shadowsocks",
		},
	}

	for _, inbound := range defaults {

		err := m.EnsureInbound(
			ctx,
			inbound.Tag,
			inbound.Port,
			inbound.Protocol,
		)

		if err != nil {
			return fmt.Errorf(
				"ensure %s: %w",
				inbound.Protocol,
				err,
			)
		}
	}

	return nil
}
