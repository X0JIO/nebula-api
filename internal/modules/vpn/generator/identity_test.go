package generator

import (
	"testing"

	"github.com/google/uuid"
)

func TestGenerateIdentity(t *testing.T) {

	id, err := GenerateIdentity()

	if err != nil {
		t.Fatal(err)
	}

	if _, err := uuid.Parse(id.UserUUID); err != nil {
		t.Fatal("invalid user uuid")
	}

	if _, err := uuid.Parse(id.SubscriptionToken); err != nil {
		t.Fatal("invalid subscription uuid")
	}

	if len(id.RealityPrivateKey) == 0 {
		t.Fatal("private key empty")
	}

	if len(id.RealityPublicKey) == 0 {
		t.Fatal("public key empty")
	}

	if len(id.RealityShortID) != 16 {
		t.Fatalf("invalid short id length: %d", len(id.RealityShortID))
	}
}
