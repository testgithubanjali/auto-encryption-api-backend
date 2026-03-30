package services

import (
	"auto-encryption-api-backend/utils"
	"log"
)

// normalize key to 32 bytes
func fixKey(key string) []byte {
	log.Println("FixKey: normalizing Key")
	k := []byte(key)

	if len(k) < 32 {
		log.Println("fixKey: Key length less than 32, padding applied")
		padding := make([]byte, 32-len(k))
		k = append(k, padding...)
	} else if len(k) > 32 {
		log.Println("fixKey: key length greater than 32, trimming applied")
		k = k[:32]
	} else {
		log.Println("fixKey: Key already 32 bytes")
	}

	return k
}

func EncryptUserText(text string, key string) (string, error) {
	log.Println("EncryptUserText: Encryption started")
	if text == "" {
		log.Println("EncryptUserText: Input text is empty")
	}

	encrypted, err := utils.EncryptText(text, fixKey(key))
	if err != nil {
		log.Println("EncryptText: Encryption failed")
		return "", err
	}
	log.Println("EncryptionText: Encryption Successful")
	return encrypted, nil
}

// 🔓 DECRYPT
func DecryptUserText(ciphertext string, key string) (string, error) {
	log.Println("Decryption started")
	if ciphertext == "" {
		log.Println("EncryptUserText: Input text is empty")
	}
	normalizedKey := fixKey(key)
	decrypted, err := utils.DecryptText(ciphertext, normalizedKey)
	if err != nil {
		log.Println("DecryptText: Decryption failed")
		return "", err
	}
	log.Println("Decryption Successful")
	return decrypted, nil
}
