package reality

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/curve25519"
)

func GenerateKeyPair() (*KeyPair, error) {
	var private [32]byte

	if _, err := rand.Read(private[:]); err != nil {
		return nil, err
	}

	// Clamp согласно X25519
	private[0] &= 248
	private[31] &= 127
	private[31] |= 64

	public, err := curve25519.X25519(
		private[:],
		curve25519.Basepoint,
	)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		PrivateKey: base64.RawURLEncoding.EncodeToString(private[:]),
		PublicKey:  base64.RawURLEncoding.EncodeToString(public),
	}, nil
}
