package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

func EncryptText(plainText string, key []byte) (string, error) {
	log.Printf("EncryptText: encrypting text of length %d with key length %d", len(plainText), len(key))

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("EncryptText: failed to create AES cipher: %v", err)
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("EncryptText: failed to create GCM: %v", err)
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Printf("EncryptText: failed to generate nonce: %v", err)
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	encoded := base64.StdEncoding.EncodeToString(cipherText)
	log.Printf("EncryptText: success, encrypted text length=%d, encoded length=%d", len(cipherText), len(encoded))
	return encoded, nil
}

func DecryptText(cipherText string, key []byte) (string, error) {
	log.Printf("DecryptText: decrypting text of length %d with key length %d", len(cipherText), len(key))

	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		log.Printf("DecryptText: failed to decode base64: %v", err)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("DecryptText: failed to create AES cipher: %v", err)
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("DecryptText: failed to create GCM: %v", err)
		return "", err
	}

	nonceSize := gcm.NonceSize()

	nonce, encryptedData := data[:nonceSize], data[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		log.Printf("DecryptText: failed to decrypt: %v", err)
		return "", err
	}

	log.Printf("DecryptText: success, decrypted text length=%d", len(plainText))
	return string(plainText), nil
}

func EncryptFile(data []byte, key []byte) ([]byte, error) {
	log.Printf("EncryptFile: encrypting file data of length %d with key length %d", len(data), len(key))

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("EncryptFile: failed to create AES cipher: %v", err)
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("EncryptFile: failed to create GCM: %v", err)
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Printf("EncryptFile: failed to generate nonce: %v", err)
		return nil, err
	}

	cipherData := gcm.Seal(nonce, nonce, data, nil)

	log.Printf("EncryptFile: success, encrypted data length=%d", len(cipherData))
	return cipherData, nil
}

func DecryptFile(cipherData []byte, key []byte) ([]byte, error) {
	log.Printf("DecryptFile: decrypting file data of length %d with key length %d", len(cipherData), len(key))

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("DecryptFile: failed to create AES cipher: %v", err)
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("DecryptFile: failed to create GCM: %v", err)
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	nonce, encryptedData := cipherData[:nonceSize], cipherData[nonceSize:]

	plainData, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		log.Printf("DecryptFile: failed to decrypt: %v", err)
		return nil, err
	}

	log.Printf("DecryptFile: success, decrypted data length=%d", len(plainData))
	return plainData, nil
}
