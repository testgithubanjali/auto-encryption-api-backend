package models

type EncodeRequest struct {
	Text string `json:"text"`
}

type DecodeRequest struct {
	Encoded string `json:"encoded"`
	Hash    string `json:"hash"` // 🔐 add this
}

type EncodeResponse struct {
	Encoded string `json:"encoded"`
}

type DecodeResponse struct {
	Text string `json:"text"`
}
