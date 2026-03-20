package handlers

import (
	"net/http"
	"os"

	"auto-encryption-api-backend/utils"
	"auto-encryption-api-backend/services"

	"github.com/gin-gonic/gin"
)

func EncryptFileHandler(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	data := make([]byte, header.Size)
	file.Read(data)

	key := []byte(os.Getenv("ENCRYPTION_KEY"))

	encryptedData, err := utils.EncryptFile(data, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
		return
	}

	path := "uploads/" + header.Filename + ".enc"

	err = services.SaveEncryptedFile(path, encryptedData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File save failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File encrypted and saved",
		"path": path,
	})
}
func DecryptFileHandler(c *gin.Context) {

	filename := c.Param("name")

	path := "uploads/" + filename

	data, err := services.ReadEncryptedFile(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File not found"})
		return
	}

	key := []byte(os.Getenv("ENCRYPTION_KEY"))

	decryptedData, err := utils.DecryptFile(data, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Decryption failed"})
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", decryptedData)
}