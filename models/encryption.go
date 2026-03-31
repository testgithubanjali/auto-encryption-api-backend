package models

type EncryptRequest struct {
	Text      string `json:"text"`
	SecretKey string `json:"secret_key"`
}
type DecryptRequest struct {
	Ciphertext string `json:"ciphertext"`
	SecretKey  string `json:"secret_key"`
	Hash       string `json:"hash"` // 🔐 add this
}
