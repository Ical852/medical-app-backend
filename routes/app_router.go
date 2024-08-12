package routes

import (
	"github.com/gin-gonic/gin"
	"medical-app-backend/config"
	"medical-app-backend/controllers"
	"medical-app-backend/middlewares"
)

func SetupRouter(userController *controllers.UserController, config *config.Config) *gin.Engine {
	router := gin.Default()

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	auth := router.Group("/auth")
	auth.Use(middlewares.AuthMiddleware(config))
	{
		// Tambahkan protected routes di sini
	}

	return router
}
