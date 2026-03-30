package models

type EntryRequest struct {
	KeyID string `json:"key_id"`
	Text  string `json:"text"`
}
