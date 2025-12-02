package routes

import (
	"banking_transaction_go/controllers"
	"banking_transaction_go/services"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	authService := &services.AuthService{}
	authController := &controllers.AuthController{Service: authService}

	g := e.Group("/auth")
	// Adding routes to the group
	g.POST("/register", authController.Register)
	g.POST("/login", authController.Login)
}