package cryptogate

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// VerifyWebhook verifies the HMAC-SHA256 signature on an incoming CryptoGate webhook.
//
// Pass the raw request body bytes, not a parsed struct.
func VerifyWebhook(secret string, rawBody []byte, signature string) bool {
	if secret == "" || len(rawBody) == 0 || signature == "" {
		return false
	}
	received := strings.TrimPrefix(signature, "sha256=")
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(rawBody)
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(received))
}
