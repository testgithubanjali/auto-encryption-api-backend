package models

// 📤 Encode Request
type EncodeRequest struct {
	Text string `json:"text"`
}

// 📥 Decode Request
type DecodeRequest struct {
	Encoded string `json:"encoded"`
}

// 📤 Encode Response
type EncodeResponse struct {
	Encoded string `json:"encoded"`
}

// 📥 Decode Response
type DecodeResponse struct {
	Text string `json:"text"`
}
