package main

import (
	"log"
	"os"

	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	log.Println("Start main server")
	log.Println("Loading environment variables")
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment")
	} else {
		log.Println(".env file loaded successfully")
	}
	log.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))

	database.ConnectDB()
	log.Println("Mongodb connected successfully")

	log.Println("Initializing gin router")
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Disable automatic redirects
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	log.Println("Router configuration completed")
	log.Println("Setting up cors")
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3005",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
			"X-Session-ID",
		},
		AllowCredentials: true,
	}))

	// Handle preflight requests
	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(200)
	})
	log.Println("cors setup completed")
	log.Println("Registering routes")
	// Register all routes
	routers.RegisterRoutes(router)
	log.Println("Routes registered successfully")

	// Read port from env
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Println("PORT not found in env, using default: 8080")
	} else {
		log.Printf("Using PORT from env: %s", port)
	}
	err = router.Run("0.0.0.0:" + port)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
