package models


type EncodeRequest struct {
	Text string `json:"text"`
}


type DecodeRequest struct {
	Encoded string `json:"encoded"`
}


type EncodeResponse struct {
	Encoded string `json:"encoded"`
}


type DecodeResponse struct {
	Text string `json:"text"`
}
