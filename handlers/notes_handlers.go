package handlers

import (
"net/http"
"os"


"auto-encryption-api-backend/models"
"auto-encryption-api-backend/services"
"auto-encryption-api-backend/utils"

"github.com/gin-gonic/gin"
"go.mongodb.org/mongo-driver/bson/primitive"


)

type NoteRequest struct {
Text string `json:"text"`
}

func CreateNote(c *gin.Context) {


var req NoteRequest

if err := c.ShouldBindJSON(&req); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Invalid request",
	})
	return
}

if req.Text == "" {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Note text cannot be empty",
	})
	return
}

userIDStr := c.MustGet("user_id").(string)

userID, err := primitive.ObjectIDFromHex(userIDStr)
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Invalid user ID",
	})
	return
}

key := []byte(os.Getenv("ENCRYPTION_KEY"))

ciphertext, err := utils.EncryptText(req.Text, key)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Encryption failed",
	})
	return
}

note := models.Note{
	UserID:     userID,
	Ciphertext: ciphertext,
}

err = services.CreateNote(note)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Failed to save note",
	})
	return
}

c.JSON(http.StatusOK, gin.H{
	"message": "Note saved",
})


}

func GetNotes(c *gin.Context) {


userIDStr := c.MustGet("user_id").(string)

userID, err := primitive.ObjectIDFromHex(userIDStr)
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Invalid user ID",
	})
	return
}

notes, err := services.GetUserNotes(userID)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Failed to fetch notes",
	})
	return
}

c.JSON(http.StatusOK, gin.H{
	"notes": notes,
})


}

func DeleteNote(c *gin.Context) {


id := c.Param("id")

noteID, err := primitive.ObjectIDFromHex(id)
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Invalid note ID",
	})
	return
}

err = services.DeleteNote(noteID)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Delete failed",
	})
	return
}

c.JSON(http.StatusOK, gin.H{
	"message": "Note deleted",
})


}
