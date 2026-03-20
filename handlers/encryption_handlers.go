package handlers

import (
	"net/http"

	"auto-encryption-api-backend/services"

	"github.com/gin-gonic/gin"
)

// 🔐 ENCRYPT
type EncryptRequest struct {
	Text      string `json:"text"`
	SecretKey string `json:"secret_key"`
}

func EncryptText(c *gin.Context) {
	var req EncryptRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Text == "" || req.SecretKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text and key required"})
		return
	}

	cipherText, err := services.EncryptUserText(req.Text, req.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ciphertext": cipherText,
	})
}

// 🔓 DECRYPT
type DecryptRequest struct {
	Ciphertext string `json:"ciphertext"`
	SecretKey  string `json:"secret_key"`
}

func DecryptText(c *gin.Context) {
	var req DecryptRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Ciphertext == "" || req.SecretKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ciphertext and key required"})
		return
	}

	plainText, err := services.DecryptUserText(req.Ciphertext, req.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Decryption failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"text": plainText, // ✅ FIXED (matches frontend)
	})
}
