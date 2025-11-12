package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func VerifySignature(payload []byte, signature, secret string) bool {
	if signature == "" {
		return false
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expected := fmt.Sprintf("sha256=%s", hex.EncodeToString(mac.Sum(nil)))
	return hmac.Equal([]byte(expected), []byte(signature))
}
