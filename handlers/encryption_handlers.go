package handlers

import (
	"log"
	"net/http"
	"time"

	"auto-encryption-api-backend/models"
	"auto-encryption-api-backend/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EncryptText(c *gin.Context) {
	log.Println("EncryptText: Starting text encryption request")

	var req models.EncryptRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("EncryptText: getting invalid request", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	log.Printf("EncryptText: Request received ,text length: %d", len(req.Text))
	if req.Text == "" || req.SecretKey == "" {
		log.Println("EncryptText: missed to get text or key")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text and key required"})
		return
	}

	log.Println("EncryptText: Encrypting Text")
	cipherText, err := services.EncryptUserText(req.Text, req.SecretKey)
	if err != nil {
		log.Println("EncryptText: Encryption failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userIDStr := c.MustGet("user_id").(string)
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	log.Printf("EncryptText: Processing for userID: %s ", userIDStr)
	if err != nil {
		log.Println("EncryptText: getting invalid user ID", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	entry := models.Entry{
		UserID:        userID,
		Type:          "encryption",
		OriginalText:  req.Text,
		ProcessedText: cipherText,
		CreatedAt:     time.Now(),
	}
	log.Println("EncryptText: Saving Encrypted Text to database")
	err = services.SaveEntry(entry)
	if err != nil {
		log.Println("EncryptText: Saving failed in Database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB save failed"})
		return
	}
	log.Println("EncryptText: Encryption completed ")
	c.JSON(http.StatusOK, gin.H{
		"ciphertext": cipherText,
	})
}

func DecryptText(c *gin.Context) {
	log.Println("Decryption Started")
	var req models.DecryptRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("DecryptText: getting invalid request", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	log.Printf("DecryptText: Request received, ciphertext length: %d", len(req.Ciphertext))
	if req.Ciphertext == "" || req.SecretKey == "" {
		log.Println("DecryptText: cipherText or secretKey missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ciphertext and key required"})
		return
	}
	log.Println("DecryptText: Decrypting text")
	plainText, err := services.DecryptUserText(req.Ciphertext, req.SecretKey)
	if err != nil {
		log.Println("DecryptText: failed to decrypt text")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userIDStr := c.MustGet("user_id").(string)
	userID, _ := primitive.ObjectIDFromHex(userIDStr)
	log.Printf("DecryptText: processing for userID: %s/n", userIDStr)
	entry := models.Entry{
		UserID:        userID,
		Type:          "decryption",
		OriginalText:  req.Ciphertext,
		ProcessedText: plainText,
		CreatedAt:     time.Now(),
	}
	log.Println("DecryptText: saving decrypting text to database")
	err = services.SaveEntry(entry)
	if err != nil {
		log.Println("DecryptText: failed to save entry", err)
	} else {
		log.Println("DecryptText: Data saved successfully")
	}
	log.Println("DecryptText: Decryption Successfull")
	c.JSON(http.StatusOK, gin.H{
		"text": plainText,
	})
}
