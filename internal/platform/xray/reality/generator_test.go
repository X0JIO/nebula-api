package reality

import "testing"

func TestGenerate(t *testing.T) {
	c, err := Generate()
	if err != nil {
		t.Fatal(err)
	}

	if c.UUID == "" {
		t.Fatal("uuid is empty")
	}

	if c.PrivateKey == "" {
		t.Fatal("private key is empty")
	}

	if c.PublicKey == "" {
		t.Fatal("public key is empty")
	}

	if c.ShortID == "" {
		t.Fatal("short id is empty")
	}
}
