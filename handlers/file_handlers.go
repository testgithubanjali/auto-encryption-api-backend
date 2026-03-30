package handlers

import (
	"auto-encryption-api-backend/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func EncryptFileHandler(c *gin.Context) {

	log.Println("Encrypting file started")

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("EncryptFileHandler: file upload failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}
	defer file.Close()

	log.Printf("EncryptFile: Received file: %s (size: %d bytes)", header.Filename, header.Size)

	data := make([]byte, header.Size)
	n, readErr := file.Read(data)
	if readErr != nil {
		log.Println("EncryptFile: File read failed", readErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File read failed"})
		return
	}

	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	if len(key) == 0 {
		log.Println("EncryptFile: server encryption key is missing")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server encryption key missing"})
		return
	}

	log.Println("EncryptFile: Encrypting file...")
	encryptedData, err := utils.EncryptFile(data[:n], key)
	if err != nil {
		log.Println("EncryptFile: encryption failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
		return
	}

	filename := header.Filename + ".enc"

	log.Println("EncryptFile: Sending encrypted file for download")

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", len(encryptedData)))

	c.Data(http.StatusOK, "application/octet-stream", encryptedData)
}

func DecryptFileHandler(c *gin.Context) {

	log.Println("File Decryption started")

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("DecryptFile: file upload failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("DecryptFile: file read failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File read failed"})
		return
	}

	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	if len(key) == 0 {
		log.Println("DecryptFile: server encryption key missing")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server encryption key missing"})
		return
	}

	decryptedData, err := utils.DecryptFile(data, key)
	if err != nil {
		log.Println("DecryptFile: decryption failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Decryption failed"})
		return
	}

	// remove .enc from filename
	filename := header.Filename
	if len(filename) > 4 && filename[len(filename)-4:] == ".enc" {
		filename = filename[:len(filename)-4]
	}

	log.Println("DecryptFile: Decryption successful")

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.Data(http.StatusOK, "application/octet-stream", decryptedData)
}
