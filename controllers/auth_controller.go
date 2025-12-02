package controllers

import (
	"net/http"
	"banking_transaction_go/services"

	"github.com/labstack/echo/v4"
)


type AuthController struct {
	Service *services.AuthService
}

func (ac *AuthController) Register(c echo.Context) error {
	var body struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	user, err := ac.Service.Register(body.Name, body.Email, body.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(201, echo.Map{
		"message": "User registered successfully",
		"user": user,
	})
}


func (ac *AuthController) Login(c echo.Context) error {
	var body struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid Request"})
	}

	access, refresh, user, err := ac.Service.Login(body.Email, body.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"access_token": access,
		"refresh_token": refresh,
		"user": user,
	})
}