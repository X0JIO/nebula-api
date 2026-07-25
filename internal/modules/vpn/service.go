package vpn

import (
	"context"

	"github.com/X0JIO/nebula-api/internal/modules/vpn/generator"
	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/X0JIO/nebula-api/internal/shared/apperrors"
)

const defaultVPNHost = "vpn.example.com"

type Service struct {
	repo *Repository
	sync *SyncService
}

func NewService(
	repo *Repository,
	sync *SyncService,
) *Service {
	return &Service{
		repo: repo,
		sync: sync,
	}
}

func (s *Service) CreateConfig(
	ctx context.Context,
	userID string,
	protocol string,
) (*CreateConfigResponse, error) {

	if protocol == "" {
		return nil, apperrors.ErrProtocolRequired
	}

	vpnUser, err := s.repo.GetVPNUser(ctx, userID)

	if err != nil {

		identity, err := generator.GenerateIdentity()
		if err != nil {
			return nil, err
		}

		vpnUser, err = s.repo.CreateVPNUser(
			ctx,
			userID,
			identity.UserUUID,
			identity.RealityPrivateKey,
			identity.RealityPublicKey,
			identity.RealityShortID,
			identity.SubscriptionToken,
		)
		if err != nil {
			return nil, err
		}

		err = s.sync.AddUser(
			ctx,
			identity.UserUUID,
			vpnUser.SubscriptionToken,
			protocol,
		)
		if err != nil {
			return nil, err
		}
	}

	identity := &generator.Identity{
		UserUUID:          vpnUser.Uuid,
		SubscriptionToken: vpnUser.SubscriptionToken,
		RealityPrivateKey: vpnUser.PrivateKey,
		RealityPublicKey:  vpnUser.PublicKey,
		RealityShortID:    vpnUser.ShortID,
	}

	var config string

	switch protocol {

	case "vless":
		config = generator.GenerateVLESS(
			identity,
			defaultVPNHost,
		)

	case "reality":
		config = generator.GenerateReality(
			identity,
			defaultVPNHost,
		)

	case "vmess":
		config = generator.GenerateVMess(
			identity,
			defaultVPNHost,
		)

	case "trojan":
		config = generator.GenerateTrojan(
			identity,
			defaultVPNHost,
		)

	case "shadowsocks":
		config = generator.GenerateShadowsocks(
			identity,
			defaultVPNHost,
		)

	default:
		return nil, apperrors.ErrUnsupportedProtocol
	}

	cfg, err := s.repo.SaveVPNConfig(
		ctx,
		db.SaveVPNConfigParams{
			VpnUserID: vpnUser.ID,
			Protocol:  protocol,
			Config:    config,
		},
	)

	if err != nil {
		return nil, err
	}

	return &CreateConfigResponse{
		ID:       cfg.ID.String(),
		Protocol: cfg.Protocol,
		Config:   cfg.Config,
	}, nil
}

func (s *Service) ListConfigs(
	ctx context.Context,
	userID string,
) (*ConfigResponse, error) {

	vpnUser, err := s.repo.GetVPNUser(ctx, userID)
	if err != nil {
		return nil, apperrors.ErrVPNUserNotFound
	}

	configs, err := s.repo.ListVPNConfigs(
		ctx,
		vpnUser.ID,
	)
	if err != nil {
		return nil, err
	}

	resp := &ConfigResponse{}

	for _, cfg := range configs {

		switch cfg.Protocol {

		case "vless":
			resp.VLESS = cfg.Config

		case "reality":
			resp.Reality = cfg.Config

		case "vmess":
			resp.VMess = cfg.Config

		case "trojan":
			resp.Trojan = cfg.Config

		case "shadowsocks":
			resp.Shadowsocks = cfg.Config
		}
	}

	return resp, nil
}

func (s *Service) Subscription(
	ctx context.Context,
	userID string,
) (string, error) {

	vpnUser, err := s.repo.GetVPNUser(ctx, userID)
	if err != nil {
		return "", err
	}

	configs, err := s.repo.ListVPNConfigs(
		ctx,
		vpnUser.ID,
	)
	if err != nil {
		return "", err
	}

	var (
		vless       string
		reality     string
		vmess       string
		trojan      string
		shadowsocks string
	)

	for _, cfg := range configs {

		switch cfg.Protocol {

		case "vless":
			vless = cfg.Config

		case "reality":
			reality = cfg.Config

		case "vmess":
			vmess = cfg.Config

		case "trojan":
			trojan = cfg.Config

		case "shadowsocks":
			shadowsocks = cfg.Config
		}
	}

	return generator.GenerateSubscription(
		vless,
		reality,
		vmess,
		trojan,
		shadowsocks,
	), nil
}

func (s *Service) SubscriptionBase64(
	ctx context.Context,
	userID string,
) (string, error) {

	vpnUser, err := s.repo.GetVPNUser(ctx, userID)
	if err != nil {
		return "", err
	}

	configs, err := s.repo.ListVPNConfigs(
		ctx,
		vpnUser.ID,
	)
	if err != nil {
		return "", err
	}

	var (
		vless       string
		reality     string
		vmess       string
		trojan      string
		shadowsocks string
	)

	for _, cfg := range configs {

		switch cfg.Protocol {

		case "vless":
			vless = cfg.Config

		case "reality":
			reality = cfg.Config

		case "vmess":
			vmess = cfg.Config

		case "trojan":
			trojan = cfg.Config

		case "shadowsocks":
			shadowsocks = cfg.Config
		}
	}

	return generator.GenerateSubscriptionBase64(
		vless,
		reality,
		vmess,
		trojan,
		shadowsocks,
	), nil
}
