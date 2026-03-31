package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashToken hashes any string using SHA-256
func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
func HashData(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
