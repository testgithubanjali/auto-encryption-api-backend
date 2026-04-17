package routers

import (
	"auto-encryption-api-backend/handlers"
	"auto-encryption-api-backend/middleware"

	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	// PUBLIC ROUTES
	router.POST("/signup", handlers.SignUpUser)
	router.POST("/login", handlers.LoginUser)
	router.POST("/refresh", handlers.RefreshToken)
	router.POST("/logout", handlers.Logout)

	// PROTECTED ROUTES
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// ✅ Create limiters (from middleware)
	encryptLimiter := middleware.EncryptLimiter()
	decryptLimiter := middleware.DecryptLimiter()
	encryptFileLimiter := middleware.EncryptFileLimiter()
	decryptFileLimiter := middleware.DecryptFileLimiter()

	{
		// user
		protected.GET("/users", handlers.UserProfile)

		// ✅ encrypt / encode
		protected.POST("/encrypt",
			tollbooth_gin.LimitHandler(encryptLimiter),
			handlers.EncryptText,
		)

		protected.POST("/encode",
			tollbooth_gin.LimitHandler(encryptLimiter),
			handlers.EncodeText,
		)

		// ✅ decrypt / decode
		protected.POST("/decrypt",
			tollbooth_gin.LimitHandler(decryptLimiter),
			handlers.DecryptText,
		)

		protected.POST("/decode",
			tollbooth_gin.LimitHandler(decryptLimiter),
			handlers.DecodeText,
		)

		// keys
		protected.POST("/keys", handlers.CreateKey)

		// entries
		protected.POST("/entries", handlers.SaveEntry)

		// ✅ file encryption
		protected.POST("/encrypt-file",
			tollbooth_gin.LimitHandler(encryptFileLimiter),
			handlers.EncryptFileHandler,
		)

		protected.POST("/decrypt-file",
			tollbooth_gin.LimitHandler(decryptFileLimiter),
			handlers.DecryptFileHandler,
		)
	}
}
