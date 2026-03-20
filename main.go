package main

import (
	"os"

	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	godotenv.Load()

	// Connect MongoDB
	database.ConnectDB()

	// Create router (without default redirect behavior)
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Disable automatic redirects
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3004",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
	}))

	// Handle preflight requests
	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(200)
	})

	// Register all routes
	routers.RegisterRoutes(router)

	// Read port from env
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// Start server
	router.Run(":" + port)
}
