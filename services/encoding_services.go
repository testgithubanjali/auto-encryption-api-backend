package services

import (
	"encoding/base64"
	"log"
)

func EncodeText(text string) (string, error) {
	log.Println("EncodeText: Encoding started")
	if text == "" {
		log.Println("EncodeText: Input text is empty")
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	log.Println("EncodeText: Encoding Successful")
	return encoded, nil
}

func DecodeText(encoded string) (string, error) {
	log.Println("DecodeText: Decoding started")

	if encoded == "" {
		log.Println("DecodeText: Input encoded string is empty")
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Println("DecodeText: Decoding failed", err)
		return "", err
	}
	log.Println("DeocodeText: Decoding Successful")
	return string(decodedBytes), nil
}
