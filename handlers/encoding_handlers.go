package handlers

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EncodeRequest struct {
	Text string `json:"text"`
}

func EncodeText(c *gin.Context) {
	var req EncodeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text required"})
		return
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(req.Text))

	c.JSON(http.StatusOK, gin.H{
		"encoded": encoded,
	})
}

type DecodeRequest struct {
	Encoded string `json:"encoded"`
}

func DecodeText(c *gin.Context) {
	var req DecodeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Encoded == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Encoded text required"})
		return
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(req.Encoded)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Decode failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"text": string(decodedBytes),
	})
}
