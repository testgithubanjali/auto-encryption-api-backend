package main
import(
	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/handlers"
	"auto-encryption-api-backend/middleware"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main(){
	database.ConnectDB()
	router := gin.Default()
router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3001"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
}))

	router.POST("/signup", handlers.SignUpUser)
	router.POST("/login", handlers.LoginUser )

	userGroup := router.Group("/users")

	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/", handlers.UserProfile)
	}
	router.Run(":8080")

	
}