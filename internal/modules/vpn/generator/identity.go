package generator

import (
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"

	"github.com/google/uuid"
)

type Identity struct {
	UserUUID string

	SubscriptionToken string

	RealityPrivateKey string
	RealityPublicKey  string
	RealityShortID    string

	ShadowsocksPassword string
}

func GenerateIdentity() (*Identity, error) {

	userUUID := uuid.New()
	subscription := uuid.New()

	curve := ecdh.X25519()

	privateKey, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.PublicKey()

	short := make([]byte, 8)
	if _, err := rand.Read(short); err != nil {
		return nil, err
	}

	ssPassword := make([]byte, 32)
	if _, err := rand.Read(ssPassword); err != nil {
		return nil, err
	}

	return &Identity{
		UserUUID: userUUID.String(),

		SubscriptionToken: subscription.String(),

		RealityPrivateKey: base64.RawURLEncoding.EncodeToString(privateKey.Bytes()),
		RealityPublicKey:  base64.RawURLEncoding.EncodeToString(publicKey.Bytes()),
		RealityShortID:    hex.EncodeToString(short),

		ShadowsocksPassword: base64.RawURLEncoding.EncodeToString(ssPassword),
	}, nil
}
