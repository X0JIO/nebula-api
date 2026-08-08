package vpn

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"testing"

	"github.com/X0JIO/nebula-api/internal/modules/vpn/generator"
	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type mockVPNRepository struct {
	getVPNUserFn func(
		context.Context,
		string,
	) (db.VpnUser, error)

	getUserEmailFn func(
		context.Context,
		string,
	) (string, error)

	createVPNUserFn func(
		context.Context,
		string,
		string,
		string,
		string,
		string,
		string,
	) (db.VpnUser, error)

	saveVPNConfigFn func(
		context.Context,
		db.SaveVPNConfigParams,
	) (db.VpnConfig, error)

	getUserEmailCalls  int
	getVPNUserCalls    int
	createVPNUserCalls int
	saveVPNConfigCalls int

	lastSaveParams db.SaveVPNConfigParams

	DeleteVPNConfigsByUserFunc func(
		ctx context.Context,
		vpnUserID pgtype.UUID,
	) error

	DeleteVPNUserFunc func(
		ctx context.Context,
		userID pgtype.UUID,
	) error
}

type mockVPNUserProvider struct{}

func (m *mockVPNUserProvider) GetByID(
	ctx context.Context,
	id string,
) (db.User, error) {

	return db.User{
		Email: "test-user@nebula.local",
	}, nil
}

func (m *mockVPNRepository) GetUserEmail(
	ctx context.Context,
	userID string,
) (string, error) {

	m.getUserEmailCalls++

	if m.getUserEmailFn != nil {
		return m.getUserEmailFn(
			ctx,
			userID,
		)
	}

	return "test-user@nebula.local", nil
}

func (m *mockVPNRepository) GetVPNUser(
	ctx context.Context,
	userID string,
) (db.VpnUser, error) {
	m.getVPNUserCalls++

	if m.getVPNUserFn != nil {
		return m.getVPNUserFn(ctx, userID)
	}

	return db.VpnUser{}, errors.New("GetVPNUser mock not configured")
}

func (m *mockVPNRepository) CreateVPNUser(
	ctx context.Context,
	userID string,
	userUUID string,
	privateKey string,
	publicKey string,
	shortID string,
	subscriptionToken string,
) (db.VpnUser, error) {
	m.createVPNUserCalls++

	if m.createVPNUserFn != nil {
		return m.createVPNUserFn(
			ctx,
			userID,
			userUUID,
			privateKey,
			publicKey,
			shortID,
			subscriptionToken,
		)
	}

	return db.VpnUser{}, errors.New("CreateVPNUser mock not configured")
}

func (m *mockVPNRepository) SaveVPNConfig(
	ctx context.Context,
	params db.SaveVPNConfigParams,
) (db.VpnConfig, error) {
	m.saveVPNConfigCalls++
	m.lastSaveParams = params

	if m.saveVPNConfigFn != nil {
		return m.saveVPNConfigFn(ctx, params)
	}

	return db.VpnConfig{}, errors.New("SaveVPNConfig mock not configured")
}

func (m *mockVPNRepository) ListVPNConfigs(
	ctx context.Context,
	vpnUserID pgtype.UUID,
) ([]db.VpnConfig, error) {
	return nil, nil
}

func (m *mockVPNRepository) DeleteVPNConfigsByUser(
	ctx context.Context,
	vpnUserID pgtype.UUID,
) error {

	if m.DeleteVPNConfigsByUserFunc != nil {
		return m.DeleteVPNConfigsByUserFunc(
			ctx,
			vpnUserID,
		)
	}

	return nil
}

func (m *mockVPNRepository) DeleteVPNUser(
	ctx context.Context,
	userID pgtype.UUID,
) error {

	if m.DeleteVPNUserFunc != nil {
		return m.DeleteVPNUserFunc(
			ctx,
			userID,
		)
	}

	return nil
}

type mockVPNServerRepository struct {
	server db.VpnServer
	err    error

	calls int
}

func (m *mockVPNServerRepository) GetActive(
	ctx context.Context,
) (db.VpnServer, error) {
	m.calls++
	return m.server, m.err
}

type mockVPNProvisioner struct {
	err error

	calls        int
	lastProtocol string
	lastUUID     string
	lastEmail    string
}

func (m *mockVPNProvisioner) Add(
	ctx context.Context,
	protocol string,
	uuid string,
	email string,
) error {
	m.calls++
	m.lastProtocol = protocol
	m.lastUUID = uuid
	m.lastEmail = email

	return m.err
}

func testIdentity() *generator.Identity {
	return &generator.Identity{
		UserUUID:            "11111111-1111-1111-1111-111111111111",
		SubscriptionToken:   "22222222-2222-2222-2222-222222222222",
		RealityPrivateKey:   "private-key",
		RealityPublicKey:    "server-public-key",
		RealityShortID:      "aabbccddeeff0011",
		ShadowsocksPassword: "shadowsocks-password",
	}
}

func testServer() generator.ServerEndpoint {
	return generator.ServerEndpoint{
		Host:      "vpn.example.com",
		Port:      8443,
		PublicKey: "server-public-key",
		ShortID:   "aabbccddeeff0011",
	}
}

func testVPNUser() db.VpnUser {
	return db.VpnUser{
		ID:                pgtype.UUID{Bytes: [16]byte{1}, Valid: true},
		Uuid:              "11111111-1111-1111-1111-111111111111",
		PrivateKey:        "private-key",
		PublicKey:         "server-public-key",
		ShortID:           "aabbccddeeff0011",
		SubscriptionToken: "22222222-2222-2222-2222-222222222222",
	}
}

func testDBServer() db.VpnServer {
	return db.VpnServer{
		Host: "vpn.example.com",
		Port: 8443,
	}
}

func TestCreateConfig(t *testing.T) {
	protocols := []string{
		"vless",
		"reality",
		"vmess",
		"trojan",
		"shadowsocks",
	}

	for _, protocol := range protocols {
		t.Run(protocol, func(t *testing.T) {
			repo := &mockVPNRepository{
				getVPNUserFn: func(
					ctx context.Context,
					userID string,
				) (db.VpnUser, error) {
					return db.VpnUser{}, errors.New("vpn user not found")
				},
				createVPNUserFn: func(
					ctx context.Context,
					userID string,
					userUUID string,
					privateKey string,
					publicKey string,
					shortID string,
					subscriptionToken string,
				) (db.VpnUser, error) {
					return testVPNUser(), nil
				},
				saveVPNConfigFn: func(
					ctx context.Context,
					params db.SaveVPNConfigParams,
				) (db.VpnConfig, error) {
					return db.VpnConfig{
						ID:       pgtype.UUID{Bytes: [16]byte{2}, Valid: true},
						Protocol: params.Protocol,
						Config:   "generated-" + params.Protocol,
					}, nil
				},
			}

			serverRepo := &mockVPNServerRepository{
				server: testDBServer(),
			}

			provisioner := &mockVPNProvisioner{}

			service := &Service{
				repo:       repo,
				serverRepo: serverRepo,
				provision:  provisioner,
				users:      &mockVPNUserProvider{},
			}

			resp, err := service.CreateConfig(
				context.Background(),
				"user-id",
				protocol,
			)

			if err != nil {
				t.Fatalf("CreateConfig() error = %v", err)
			}

			if resp == nil {
				t.Fatal("CreateConfig() returned nil response")
			}

			if resp.Protocol != protocol {
				t.Fatalf(
					"unexpected protocol: want %q, got %q",
					protocol,
					resp.Protocol,
				)
			}

			if resp.Config == "" {
				t.Fatal("CreateConfig() returned empty config")
			}

			if serverRepo.calls != 1 {
				t.Fatalf(
					"GetActive() calls: want 1, got %d",
					serverRepo.calls,
				)
			}

			if repo.createVPNUserCalls != 1 {
				t.Fatalf(
					"CreateVPNUser() calls: want 1, got %d",
					repo.createVPNUserCalls,
				)
			}

			if provisioner.calls != 1 {
				t.Fatalf(
					"provision.Add() calls: want 1, got %d",
					provisioner.calls,
				)
			}

			if provisioner.lastProtocol != protocol {
				t.Fatalf(
					"provision protocol: want %q, got %q",
					protocol,
					provisioner.lastProtocol,
				)
			}

			if repo.saveVPNConfigCalls != 1 {
				t.Fatalf(
					"SaveVPNConfig() calls: want 1, got %d",
					repo.saveVPNConfigCalls,
				)
			}

			if repo.lastSaveParams.Protocol != protocol {
				t.Fatalf(
					"saved protocol: want %q, got %q",
					protocol,
					repo.lastSaveParams.Protocol,
				)
			}
		})
	}
}

func TestGenerateVLESS(t *testing.T) {
	identity := testIdentity()
	server := testServer()

	got := generator.GenerateVLESS(identity, server)

	if !strings.HasPrefix(got, "vless://") {
		t.Fatalf("invalid VLESS prefix: %s", got)
	}

	if !strings.Contains(got, identity.UserUUID) {
		t.Fatal("VLESS config does not contain user UUID")
	}

	if !strings.Contains(got, "vpn.example.com:8443") {
		t.Fatal("VLESS config does not contain server endpoint")
	}

	if !strings.Contains(got, "security=reality") {
		t.Fatal("VLESS config does not use Reality")
	}

	if !strings.Contains(got, "flow=xtls-rprx-vision") {
		t.Fatal("VLESS config does not contain expected flow")
	}

	if !strings.Contains(got, "sid="+server.ShortID) {
		t.Fatal("VLESS config does not contain short ID")
	}

	if !strings.Contains(got, "sni=vpn.example.com") {
		t.Fatal("VLESS config does not contain SNI")
	}

	if !strings.Contains(got, "pbk="+url.QueryEscape(server.PublicKey)) {
		t.Fatal("VLESS config does not contain public key")
	}
}

func TestGenerateReality(t *testing.T) {
	identity := testIdentity()
	server := testServer()

	got := generator.GenerateReality(identity, server)

	if !strings.HasPrefix(got, "vless://") {
		t.Fatalf("invalid Reality prefix: %s", got)
	}

	if !strings.Contains(got, identity.UserUUID) {
		t.Fatal("Reality config does not contain user UUID")
	}

	if !strings.Contains(got, "vpn.example.com:8443") {
		t.Fatal("Reality config does not contain server endpoint")
	}

	if !strings.Contains(got, "security=reality") {
		t.Fatal("Reality config does not use Reality security")
	}

	if !strings.Contains(got, "sid="+server.ShortID) {
		t.Fatal("Reality config does not contain short ID")
	}

	if !strings.Contains(got, "sni=vpn.example.com") {
		t.Fatal("Reality config does not contain SNI")
	}
}

func TestGenerateVMess(t *testing.T) {
	identity := testIdentity()
	server := testServer()

	got := generator.GenerateVMess(identity, server)

	if !strings.HasPrefix(got, "vmess://") {
		t.Fatalf("invalid VMess prefix: %s", got)
	}

	raw := strings.TrimPrefix(got, "vmess://")

	data, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		t.Fatalf("failed to decode VMess config: %v", err)
	}

	var cfg map[string]string

	if err := json.Unmarshal(data, &cfg); err != nil {
		t.Fatalf("failed to parse VMess JSON: %v", err)
	}

	if cfg["v"] != "2" {
		t.Fatalf("unexpected VMess version: %s", cfg["v"])
	}

	if cfg["add"] != server.Host {
		t.Fatalf("unexpected VMess host: %s", cfg["add"])
	}

	if cfg["port"] != "8443" {
		t.Fatalf("unexpected VMess port: %s", cfg["port"])
	}

	if cfg["id"] != identity.UserUUID {
		t.Fatalf("unexpected VMess UUID: %s", cfg["id"])
	}

	if cfg["net"] != "tcp" {
		t.Fatalf("unexpected VMess network: %s", cfg["net"])
	}

	if cfg["tls"] != "tls" {
		t.Fatalf("unexpected VMess TLS: %s", cfg["tls"])
	}
}

func TestGenerateTrojan(t *testing.T) {
	identity := testIdentity()
	server := testServer()

	got := generator.GenerateTrojan(identity, server)

	if !strings.HasPrefix(got, "trojan://") {
		t.Fatalf("invalid Trojan prefix: %s", got)
	}

	if !strings.Contains(got, identity.UserUUID) {
		t.Fatal("Trojan config does not contain UUID")
	}

	if !strings.Contains(got, "vpn.example.com:8443") {
		t.Fatal("Trojan config does not contain server endpoint")
	}

	if !strings.Contains(got, "security=tls") {
		t.Fatal("Trojan config does not use TLS")
	}

	if !strings.Contains(got, "sni=vpn.example.com") {
		t.Fatal("Trojan config does not contain SNI")
	}
}

func TestGenerateShadowsocks(t *testing.T) {
	identity := testIdentity()
	server := testServer()

	got := generator.GenerateShadowsocks(identity, server)

	if !strings.HasPrefix(got, "ss://") {
		t.Fatalf("invalid Shadowsocks prefix: %s", got)
	}

	if !strings.Contains(got, "chacha20-ietf-poly1305") {
		t.Fatal("Shadowsocks config does not contain expected method")
	}

	if !strings.Contains(got, identity.ShadowsocksPassword) {
		t.Fatal("Shadowsocks config does not contain password")
	}

	if !strings.Contains(got, "vpn.example.com:8443") {
		t.Fatal("Shadowsocks config does not contain server endpoint")
	}
}

func TestGenerateSubscription(t *testing.T) {
	got := generator.GenerateSubscription(
		"vless://test",
		"",
		"vmess://test",
		"",
		"ss://test",
	)

	want := "vless://test\nvmess://test\nss://test"

	if got != want {
		t.Fatalf("unexpected subscription:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestGenerateSubscriptionBase64(t *testing.T) {
	got := generator.GenerateSubscriptionBase64(
		"vless://test",
		"",
		"vmess://test",
		"",
		"ss://test",
	)

	decoded, err := base64.StdEncoding.DecodeString(got)
	if err != nil {
		t.Fatalf("failed to decode subscription: %v", err)
	}

	want := "vless://test\nvmess://test\nss://test"

	if string(decoded) != want {
		t.Fatalf("unexpected decoded subscription:\nwant: %q\ngot:  %q", want, string(decoded))
	}
}

func TestCreateConfigExistingVPNUser(t *testing.T) {
	repo := &mockVPNRepository{
		getVPNUserFn: func(
			ctx context.Context,
			userID string,
		) (db.VpnUser, error) {
			return testVPNUser(), nil
		},
		saveVPNConfigFn: func(
			ctx context.Context,
			params db.SaveVPNConfigParams,
		) (db.VpnConfig, error) {
			return db.VpnConfig{
				ID:       pgtype.UUID{Bytes: [16]byte{2}, Valid: true},
				Protocol: params.Protocol,
				Config:   "existing-user-config",
			}, nil
		},
	}

	serverRepo := &mockVPNServerRepository{
		server: testDBServer(),
	}

	provisioner := &mockVPNProvisioner{}

	service := &Service{
		repo:       repo,
		serverRepo: serverRepo,
		provision:  provisioner,
		users:      &mockVPNUserProvider{},
	}

	resp, err := service.CreateConfig(
		context.Background(),
		"user-id",
		"vless",
	)

	if err != nil {
		t.Fatalf("CreateConfig() error = %v", err)
	}

	if resp == nil {
		t.Fatal("CreateConfig() returned nil response")
	}

	if repo.createVPNUserCalls != 0 {
		t.Fatalf(
			"CreateVPNUser() calls: want 0, got %d",
			repo.createVPNUserCalls,
		)
	}

	if provisioner.calls != 0 {
		t.Fatalf(
			"provision.Add() calls: want 0, got %d",
			provisioner.calls,
		)
	}

	if repo.saveVPNConfigCalls != 1 {
		t.Fatalf(
			"SaveVPNConfig() calls: want 1, got %d",
			repo.saveVPNConfigCalls,
		)
	}
}

func TestCreateConfigGetActiveServerError(t *testing.T) {
	expectedErr := errors.New("database unavailable")

	repo := &mockVPNRepository{
		getVPNUserFn: func(
			ctx context.Context,
			userID string,
		) (db.VpnUser, error) {
			return testVPNUser(), nil
		},
	}

	serverRepo := &mockVPNServerRepository{
		err: expectedErr,
	}

	provisioner := &mockVPNProvisioner{}

	service := &Service{
		repo:       repo,
		serverRepo: serverRepo,
		provision:  provisioner,
		users:      &mockVPNUserProvider{},
	}

	_, err := service.CreateConfig(
		context.Background(),
		"user-id",
		"vless",
	)

	if !errors.Is(err, expectedErr) {
		t.Fatalf(
			"expected GetActive() error %v, got %v",
			expectedErr,
			err,
		)
	}

	if provisioner.calls != 0 {
		t.Fatalf(
			"provision.Add() calls: want 0, got %d",
			provisioner.calls,
		)
	}

	if repo.saveVPNConfigCalls != 0 {
		t.Fatalf(
			"SaveVPNConfig() calls: want 0, got %d",
			repo.saveVPNConfigCalls,
		)
	}
}

func TestCreateConfigProvisionError(t *testing.T) {
	expectedErr := errors.New("xray unavailable")

	repo := &mockVPNRepository{
		getVPNUserFn: func(
			ctx context.Context,
			userID string,
		) (db.VpnUser, error) {
			return db.VpnUser{}, errors.New("vpn user not found")
		},
		createVPNUserFn: func(
			ctx context.Context,
			userID string,
			userUUID string,
			privateKey string,
			publicKey string,
			shortID string,
			subscriptionToken string,
		) (db.VpnUser, error) {
			return testVPNUser(), nil
		},
		saveVPNConfigFn: func(
			ctx context.Context,
			params db.SaveVPNConfigParams,
		) (db.VpnConfig, error) {
			t.Fatal("SaveVPNConfig() must not be called after provision failure")
			return db.VpnConfig{}, nil
		},
	}

	serverRepo := &mockVPNServerRepository{
		server: testDBServer(),
	}

	provisioner := &mockVPNProvisioner{
		err: expectedErr,
	}

	service := &Service{
		repo:       repo,
		serverRepo: serverRepo,
		provision:  provisioner,
		users:      &mockVPNUserProvider{},
	}

	_, err := service.CreateConfig(
		context.Background(),
		"user-id",
		"vless",
	)

	if !errors.Is(err, expectedErr) {
		t.Fatalf(
			"expected provision error %v, got %v",
			expectedErr,
			err,
		)
	}

	if provisioner.calls != 1 {
		t.Fatalf(
			"provision.Add() calls: want 1, got %d",
			provisioner.calls,
		)
	}

	if repo.saveVPNConfigCalls != 0 {
		t.Fatalf(
			"SaveVPNConfig() calls: want 0, got %d",
			repo.saveVPNConfigCalls,
		)
	}
}

func TestCreateConfigCreateVPNUserError(t *testing.T) {
	expectedErr := errors.New("failed to create vpn user")

	repo := &mockVPNRepository{
		getVPNUserFn: func(
			ctx context.Context,
			userID string,
		) (db.VpnUser, error) {
			return db.VpnUser{}, errors.New("vpn user not found")
		},
		createVPNUserFn: func(
			ctx context.Context,
			userID string,
			userUUID string,
			privateKey string,
			publicKey string,
			shortID string,
			subscriptionToken string,
		) (db.VpnUser, error) {
			return db.VpnUser{}, expectedErr
		},
	}

	serverRepo := &mockVPNServerRepository{
		server: testDBServer(),
	}

	provisioner := &mockVPNProvisioner{}

	service := &Service{
		repo:       repo,
		serverRepo: serverRepo,
		provision:  provisioner,
		users:      &mockVPNUserProvider{},
	}

	_, err := service.CreateConfig(
		context.Background(),
		"user-id",
		"vless",
	)

	if !errors.Is(err, expectedErr) {
		t.Fatalf(
			"expected CreateVPNUser error %v, got %v",
			expectedErr,
			err,
		)
	}

	if repo.createVPNUserCalls != 1 {
		t.Fatalf(
			"CreateVPNUser() calls: want 1, got %d",
			repo.createVPNUserCalls,
		)
	}

	if provisioner.calls != 0 {
		t.Fatalf(
			"provision.Add() calls: want 0, got %d",
			provisioner.calls,
		)
	}

	if repo.saveVPNConfigCalls != 0 {
		t.Fatalf(
			"SaveVPNConfig() calls: want 0, got %d",
			repo.saveVPNConfigCalls,
		)
	}
}

func TestCreateConfigSaveVPNConfigError(t *testing.T) {
	expectedErr := errors.New("failed to save vpn config")

	repo := &mockVPNRepository{
		getVPNUserFn: func(
			ctx context.Context,
			userID string,
		) (db.VpnUser, error) {
			return db.VpnUser{}, errors.New("vpn user not found")
		},
		createVPNUserFn: func(
			ctx context.Context,
			userID string,
			userUUID string,
			privateKey string,
			publicKey string,
			shortID string,
			subscriptionToken string,
		) (db.VpnUser, error) {
			return testVPNUser(), nil
		},
		saveVPNConfigFn: func(
			ctx context.Context,
			params db.SaveVPNConfigParams,
		) (db.VpnConfig, error) {
			return db.VpnConfig{}, expectedErr
		},
	}
	serverRepo := &mockVPNServerRepository{
		server: testDBServer(),
	}

	provisioner := &mockVPNProvisioner{}

	service := &Service{
		repo:       repo,
		serverRepo: serverRepo,
		provision:  provisioner,
		users:      &mockVPNUserProvider{},
	}

	_, err := service.CreateConfig(
		context.Background(),
		"user-id",
		"vless",
	)

	if !errors.Is(err, expectedErr) {
		t.Fatalf(
			"expected SaveVPNConfig error %v, got %v",
			expectedErr,
			err,
		)
	}

	if provisioner.calls != 1 {
		t.Fatalf(
			"provision.Add() calls: want 1, got %d",
			provisioner.calls,
		)
	}

	if repo.saveVPNConfigCalls != 1 {
		t.Fatalf(
			"SaveVPNConfig() calls: want 1, got %d",
			repo.saveVPNConfigCalls,
		)
	}
}
