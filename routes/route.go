package routes

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo) {

	api := e.Group("/auth")
	// Adding routes to the group
	api.GET("/register")
}