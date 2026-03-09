package routers
import (
	"auto-encryption-api-backend/handlers"
	"auto-encryption-api-backend/middleware"

	"github.com/gin-gonic/gin"
)
func RegisterRoutes(router *gin.Engine){
	router.POST("/signup", handlers.SignUpUser)
	router.POST("/login", handlers.LoginUser )
		userGroup := router.Group("/users")

	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/", handlers.UserProfile)
	}
}
