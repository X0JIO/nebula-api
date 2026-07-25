package reality

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateShortID() (string, error) {
	buf := make([]byte, 8) // 16 hex символов

	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	return hex.EncodeToString(buf), nil
}
