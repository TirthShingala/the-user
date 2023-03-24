package routes

import (
	controller "github.com/TirthShingala/the-user/controllers"
	"github.com/TirthShingala/the-user/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	userRoutes := routes.Group("/users")

	userRoutes.Use(middleware.Authenticate())

	userRoutes.GET("/", controller.GetUsers())
	userRoutes.GET("/profile", controller.GetUser())
}
