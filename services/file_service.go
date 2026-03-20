package services

import (
	"os"
)

func SaveEncryptedFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func ReadEncryptedFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}