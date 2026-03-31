package handlers

import (
	"auto-encryption-api-backend/models"
	"auto-encryption-api-backend/services"
	"auto-encryption-api-backend/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *gin.Context) {
	log.Println("SignUpUser: Starting user signup request")

	var req models.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Signup: invalid request from the user: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	log.Printf("SignupUser: Request received for email: %s\n", req.Email)

	// Check if user already exists
	existingUser, _ := services.GetUserByEmail(req.Email)
	if existingUser != nil {
		log.Println("Error checking existing user ", req.Email)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Println("error during password hashing ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password hashing failed",
		})
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	log.Println("SignupUser: Creating user in database")
	err = services.CreateUser(user)
	if err != nil {
		log.Println("error during user creation ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User creation failed",
		})
		return
	}

	log.Printf("SignUpUser: user created successfull for email %s\n", req.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successful",
	})
}

func LoginUser(c *gin.Context) {
	log.Println("LoginUser: starting user login request")
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("LoginUser: failed to get login request", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	log.Printf("LoginUser: Login attempt for email: %s\n", req.Email)
	user, err := services.GetUserByEmail(req.Email)
	if err != nil || user == nil {
		log.Println("LoginUser: error during getting email", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	log.Printf("LoginUser: comparing entered password with hashed password")
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Println("LoginUser: passsword mismatched", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	log.Println("LoginUser: generating tokens")
	accessToken, err := utils.GenerateAccessToken(user.ID.Hex())
	if err != nil {
		log.Println("LoginUser: failed to generate access token", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate access token",
		})
		return
	}
	log.Println("LoginUser: generating tokens")
	refreshToken, err := utils.GenerateRefreshToken(user.ID.Hex())

	if err != nil {
		log.Println("failed to generate refresh token", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate refresh token",
		})
		return
	}
	log.Println("creating sessions")
	sessionID := uuid.New().String()
	// 🔐 Hash refresh token before storing
	hashedRefreshToken := utils.HashToken(refreshToken)

	err = services.CreateSession(user.ID.Hex(), sessionID, hashedRefreshToken)
	if err != nil {
		log.Println("Session creation failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "session creation failed",
		})
		return
	}
	log.Printf("LoginUser: User loggin in successfully for userID: %s\n", user.ID.Hex())
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"session_id":    sessionID,
	})
}
func UserProfile(c *gin.Context) {
	log.Println("userprofile api called")

	userID := c.MustGet("user_id")
	log.Printf("UserProfile: fetching profile for useraID: %s\n", userID)
	userID2, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		log.Println("Userprofile: invalid userID", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	user, err := services.GetUserByID(userID2)
	if err != nil {
		log.Println("UserProfile: failed to get user", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User not found",
		})
		return
	}

	log.Printf("UserProfile: user profile successfully loaded for email: %s\n", user.Email)
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.ID.Hex(),
			"email": user.Email,
		},
	})
}
func Logout(c *gin.Context) {
	log.Println("Logout API called")

	sessionID := c.GetHeader("Session-ID")

	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Session ID required",
		})
		return
	}

	err := services.ExpireSession(sessionID)
	if err != nil {
		log.Println("Logout failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Logout failed",
		})
		return
	}

	log.Println("Logout completed successfully")

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
