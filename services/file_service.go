package services

import (
	"log"
	"os"
)

func SaveEncryptedFile(path string, data []byte) error {
	log.Printf("SaveEncryptedFile: saving file at path: %s", path)
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		log.Println("SaveEncryptedFile: write failed: ", err)
		return err
	}
	log.Println("SaveEncryptedFile: File Saved Successfully")
	return nil
}

func ReadEncryptedFile(path string) ([]byte, error) {
	log.Printf("ReadEnccryptedFile: reading file from path: %s", path)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Println("ReadEncryptedFile: Read failed: ", err)
		return nil, err
	}
	log.Printf("ReadEncryptedFile: File read successfully, size: %d bytes", len(data))
	return data, nil
}
