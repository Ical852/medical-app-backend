package main

import (
	"medical-app-backend/config"
	"medical-app-backend/controllers"
	"medical-app-backend/repositories"
	"medical-app-backend/routes"
	"medical-app-backend/services"
)

func main() {
	cfg := config.NewConfig()
	defer cfg.DB.Close()

	userRepository := repositories.NewUserRepository(cfg.DB)
	userService := services.NewUserService(userRepository, cfg.JWTSecret)
	userController := controllers.NewUserController(userService)

	router := routes.SetupRouter(userController, cfg)
	router.Run(":8080")
}
