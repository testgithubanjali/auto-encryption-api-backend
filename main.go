package main
import(
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.POST("/signup", SignUpUser)
	router.POST("/login", LoginUser )

	userGroup := router.Group("/users")

	userGroup.Use(AuthMiddleware())
	{
		userGroup.GET("/", UserProfile)
	}
	router.Run(":8080")

	
}