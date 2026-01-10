package go_kit

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
)

func Password(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+[];,./:"
	password := make([]byte, length)
	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[n.Int64()]
	}
	return string(password), nil
}

func OpaqueToken(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	// URL-safe, no padding
	return base64.RawURLEncoding.EncodeToString(b)
}

func SHA256(bytes int) string {
	randomBytes := make([]byte, bytes)
	if _, err := rand.Read(randomBytes); err != nil {
		return ""
	}

	rawCode := base64.RawURLEncoding.EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(rawCode))
	codeHash := fmt.Sprintf("%x", hash[:])

	return codeHash
}
