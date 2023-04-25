package routes

import (
	controller "github.com/TirthShingala/the-user/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {
	authRoutes := routes.Group("/auth")

	authRoutes.POST("/signup", controller.Signup())
	authRoutes.POST("/login", controller.Login())
	authRoutes.POST("/google", controller.Google())
	authRoutes.GET("/token", controller.Token())
}
