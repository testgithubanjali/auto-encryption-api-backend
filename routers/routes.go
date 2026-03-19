package routers

import (
	"auto-encryption-api-backend/handlers"
	"auto-encryption-api-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	// PUBLIC ROUTES
	router.POST("/signup", handlers.SignUpUser)
	router.POST("/login", handlers.LoginUser)
	router.POST("/refresh", handlers.RefreshToken)

	// PROTECTED ROUTES
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	{
		// user
		protected.GET("/users", handlers.UserProfile)

		// encryption
		protected.POST("/encrypt", handlers.EncryptText)
		protected.POST("/decrypt", handlers.DecryptText)

		// keys
		protected.POST("/keys", handlers.CreateKey)

		// entries
		protected.POST("/entries", handlers.SaveEntry)

		// notes
		protected.POST("/notes", handlers.CreateNote)
		protected.GET("/notes", handlers.GetNotes)
		protected.DELETE("/notes/:id", handlers.DeleteNote)

		// file encryption
		protected.POST("/encrypt-file", handlers.EncryptFileHandler)
		protected.GET("/decrypt-file/:name", handlers.DecryptFileHandler)
		}
	}
