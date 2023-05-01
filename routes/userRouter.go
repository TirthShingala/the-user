package routes

import (
	controller "github.com/TirthShingala/the-user/controllers"
	"github.com/TirthShingala/the-user/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	userRoutes := routes.Group("/user")

	userRoutes.Use(middleware.Authenticate())

	userRoutes.GET("/", controller.GetUsers())
	userRoutes.GET("/profile", controller.GetUser())
	userRoutes.POST("/profile/upload-url", controller.UploadURL())
}
