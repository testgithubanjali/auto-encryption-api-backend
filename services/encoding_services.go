package services

import (
	"encoding/base64"
)

// 🔐 ENCODE SERVICE
func EncodeText(text string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	return encoded, nil
}

// 🔓 DECODE SERVICE
func DecodeText(encoded string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
