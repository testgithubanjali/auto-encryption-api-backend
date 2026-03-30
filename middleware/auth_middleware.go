package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"auto-encryption-api-backend/services"
	"auto-encryption-api-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Printf("AuthMiddleware: Processing request → %s", c.FullPath())

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("AuthMiddleware: Missing Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		log.Println("AuthMiddleware: Authorization header received")

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Printf("AuthMiddleware: Invalid auth format → %v", parts)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		log.Println("AuthMiddleware: Token extracted successfully")

		claims := jwt.MapClaims{}

		// 🔹 3. Parse JWT
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

			log.Println("AuthMiddleware: Validating JWT signing method")

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("AuthMiddleware: Unexpected signing method")
				return nil, fmt.Errorf("unexpected signing method")
			}

			return utils.GetJWTSecret(), nil
		})

		// 🔥 STRICT VALIDATION
		if err != nil {
			log.Println("AuthMiddleware: JWT parse error →", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if token == nil || !token.Valid {
			log.Println("AuthMiddleware: Token is invalid")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		log.Println("AuthMiddleware: Token validated successfully")

		// 🔹 4. Check expiration manually
		exp, ok := claims["exp"].(float64)
		if !ok {
			log.Println("AuthMiddleware: exp claim missing")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if int64(exp) < time.Now().Unix() {
			log.Println("AuthMiddleware: Token expired")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		// 🔹 5. Extract user_id
		userID, ok := claims["user_id"].(string)
		if !ok {
			log.Println("AuthMiddleware: user_id missing in token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID missing in token"})
			c.Abort()
			return
		}

		log.Printf("AuthMiddleware: user_id extracted → %s", userID)

		sessionID := c.GetHeader("X-Session-ID")
		if sessionID == "" {
			log.Println("AuthMiddleware: Session ID missing")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session ID missing"})
			c.Abort()
			return
		}

		log.Printf("AuthMiddleware: Session ID received → %s", sessionID)

		session, err := services.GetSessionByID(sessionID)
		if err != nil || session == nil {
			log.Println("AuthMiddleware: Session not found in DB:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session invalid"})
			c.Abort()
			return
		}

		log.Println("AuthMiddleware: Session found in DB")

		if session.IsExpired {
			log.Printf("AuthMiddleware: Session expired → %s", sessionID)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
			c.Abort()
			return
		}

		log.Println("AuthMiddleware: Session is active")

		if session.UserID.Hex() != userID {
			log.Printf("AuthMiddleware: Session mismatch (sessionUser=%s, tokenUser=%s)", session.UserID.Hex(), userID)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session mismatch"})
			c.Abort()
			return
		}

		log.Println("AuthMiddleware: Session matches user")

		log.Printf("AuthMiddleware: Authentication SUCCESS → user=%s, path=%s", userID, c.FullPath())

		c.Set("user_id", userID)

		c.Next()
	}
}
