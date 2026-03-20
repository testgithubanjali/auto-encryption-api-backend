package services

import "auto-encryption-api-backend/utils"

// normalize key to 32 bytes
func fixKey(key string) []byte {
	k := []byte(key)

	if len(k) < 32 {
		padding := make([]byte, 32-len(k))
		k = append(k, padding...)
	} else if len(k) > 32 {
		k = k[:32]
	}

	return k
}

// 🔐 ENCRYPT
func EncryptUserText(text string, key string) (string, error) {
	return utils.EncryptText(text, fixKey(key))
}

// 🔓 DECRYPT
func DecryptUserText(ciphertext string, key string) (string, error) {
	return utils.DecryptText(ciphertext, fixKey(key))
}
