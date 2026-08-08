package vpn

import (
	"context"
	"encoding/base64"

	"github.com/X0JIO/nebula-api/internal/modules/vpn/generator"
	"github.com/X0JIO/nebula-api/internal/modules/vpn/provision"
	"github.com/X0JIO/nebula-api/internal/modules/vpn/server"
	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/X0JIO/nebula-api/internal/shared/apperrors"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repo       vpnRepository
	sync       *SyncService
	provision  vpnProvisioner
	serverRepo vpnServerRepository
	users      vpnUserProvider
}

type vpnRepository interface {
	GetVPNUser(ctx context.Context, userID string) (db.VpnUser, error)

	CreateVPNUser(
		ctx context.Context,
		userID string,
		userUUID string,
		privateKey string,
		publicKey string,
		shortID string,
		subscriptionToken string,
	) (db.VpnUser, error)

	SaveVPNConfig(
		ctx context.Context,
		params db.SaveVPNConfigParams,
	) (db.VpnConfig, error)

	ListVPNConfigs(
		ctx context.Context,
		vpnUserID pgtype.UUID,
	) ([]db.VpnConfig, error)

	DeleteVPNConfigsByUser(
		ctx context.Context,
		vpnUserID pgtype.UUID,
	) error

	DeleteVPNUser(
		ctx context.Context,
		userID pgtype.UUID,
	) error

	GetVPNUserBySubscription(
		ctx context.Context,
		token string,
	) (db.VpnUser, error)
}

type vpnServerRepository interface {
	GetActive(ctx context.Context) (db.VpnServer, error)
}

type vpnUserProvider interface {
	GetByID(
		ctx context.Context,
		id string,
	) (db.User, error)
}

type vpnProvisioner interface {
	Add(
		ctx context.Context,
		protocol string,
		uuid string,
		email string,
	) error
}

func NewService(
	repo *Repository,
	sync *SyncService,
	provision *provision.Service,
	serverRepo *server.Repository,
	users vpnUserProvider,
) *Service {
	return &Service{
		repo:       repo,
		sync:       sync,
		provision:  provision,
		serverRepo: serverRepo,
		users:      users,
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

	vpnServer, err := s.serverRepo.GetActive(ctx)

	if err != nil {
		return nil, err
	}

	var email string

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

		user, err := s.users.GetByID(
			ctx,
			userID,
		)

		if err != nil {
			return nil, err
		}

		email = user.Email

		err = s.provision.Add(
			ctx,
			protocol,
			identity.UserUUID,
			email,
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

	serverEndpoint := generator.ServerEndpoint{
		Host: vpnServer.Host,
		Port: int(vpnServer.Port),
	}

	if vpnServer.PublicKey.Valid {
		serverEndpoint.PublicKey = vpnServer.PublicKey.String
	}

	if vpnServer.ShortID.Valid {
		serverEndpoint.ShortID = vpnServer.ShortID.String
	}

	switch protocol {

	case "vless":
		config = generator.GenerateVLESS(
			identity,
			serverEndpoint,
		)

	case "reality":
		config = generator.GenerateReality(
			identity,
			serverEndpoint,
		)

	case "vmess":
		config = generator.GenerateVMess(
			identity,
			serverEndpoint,
		)

	case "trojan":
		config = generator.GenerateTrojan(
			identity,
			serverEndpoint,
		)

	case "shadowsocks":
		config = generator.GenerateShadowsocks(
			identity,
			serverEndpoint,
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

func (s *Service) PublicSubscription(
	ctx context.Context,
	token string,
) (string, error) {

	vpnUser, err := s.repo.GetVPNUserBySubscription(
		ctx,
		token,
	)

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

func (s *Service) PublicSubscriptionBase64(
	ctx context.Context,
	token string,
) (string, error) {

	sub, err := s.PublicSubscription(
		ctx,
		token,
	)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(
		[]byte(sub),
	), nil
}

func (s *Service) DeleteAccount(
	ctx context.Context,
	userID string,
) error {

	vpnUser, err := s.repo.GetVPNUser(
		ctx,
		userID,
	)

	if err != nil {
		return apperrors.ErrVPNUserNotFound
	}

	err = s.repo.DeleteVPNConfigsByUser(
		ctx,
		vpnUser.ID,
	)

	if err != nil {
		return err
	}

	userUUID, err := parseUUID(userID)

	if err != nil {
		return err
	}

	err = s.repo.DeleteVPNUser(
		ctx,
		userUUID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAccount(
	ctx context.Context,
	userID string,
) (*VPNAccountResponse, error) {

	vpnUser, err := s.repo.GetVPNUser(
		ctx,
		userID,
	)

	if err != nil {
		return nil, apperrors.ErrVPNUserNotFound
	}

	return &VPNAccountResponse{
		ID:                vpnUser.ID.String(),
		UUID:              vpnUser.Uuid,
		SubscriptionToken: vpnUser.SubscriptionToken,
		PublicKey:         vpnUser.PublicKey,
		ShortID:           vpnUser.ShortID,
		CreatedAt:         vpnUser.CreatedAt.Time,
	}, nil
}
