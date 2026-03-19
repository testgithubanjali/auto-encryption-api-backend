package handlers

import (
	"net/http"

	"auto-encryption-api-backend/services"

	"github.com/gin-gonic/gin"
)
type EncryptRequest struct {
	Text string `json:"text"`
}
func EncryptText(c *gin.Context){
	var req EncryptRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	cipherText, err := services.EncryptUserText(req.Text)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
		return 
	}
	c.JSON(http.StatusOK, gin.H{
		"ciphertext": cipherText,
	})
}
type DecryptRequest struct {
	Ciphertext string `json:"ciphertext"`
}
func DecryptText(c *gin.Context){
	var req DecryptRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request"})
		return
	}
	plainText, err := services.DecryptUserText(req.Ciphertext)
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Decryption failed"})
					return
	}
	c.JSON(http.StatusOK, gin.H{
		"plaintext": plainText,
	})
	
}

