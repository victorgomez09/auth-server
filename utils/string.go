package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// randString generates a random string of length n and returns its
// base64-encoded version.
func RandString(n int) (string, error) {
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}
