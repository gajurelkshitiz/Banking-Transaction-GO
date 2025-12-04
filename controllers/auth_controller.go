package controllers

import (
	"banking_transaction_go/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Service *services.AuthService
}

func (ac *AuthController) Register(c echo.Context) error {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		// return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
		return JSONError(c, http.StatusBadRequest, "", "Invalid request")
	}

	user, err := ac.Service.Register(body.Name, body.Email, body.Password)
	if err != nil {
		// return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		return JSONError(c, http.StatusBadRequest, err.Error(), "")
	}

	// return c.JSON(201, echo.Map{
	// 	"message": "User registered successfully",
	// 	"user":    user,
	// })
	return JSONSuccess(c, http.StatusCreated, map[string]interface{}{"user": user}, "User registered successfully")
}

func (ac *AuthController) Login(c echo.Context) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		// return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid Request"})
		JSONError(c, http.StatusBadRequest, "", "Invalid request")
	}

	access, refresh, user, err := ac.Service.Login(body.Email, body.Password)
	if err != nil {
		// return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
		return JSONError(c, http.StatusUnauthorized, err.Error(), "")
	}

	// return c.JSON(200, echo.Map{
	// 	"access_token":  access,
	// 	"refresh_token": refresh,
	// 	"user":          user,
	// })
	return JSONSuccess(c, http.StatusOK, map[string]interface{}{
		"access_token": access,
		"refresh_token": refresh,
		"user": user,
	}, "")
}

func (ac *AuthController) Refresh(c echo.Context) error {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.Bind(&body); err != nil {
		// return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
		return JSONError(c, http.StatusBadRequest, "", "Invalid request")
	}

	accessToken, err := ac.Service.Refresh(body.RefreshToken)
	if err != nil {
		// return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
		return JSONError(c, http.StatusUnauthorized, err.Error(), "")
	}

	// return c.JSON(http.StatusOK, echo.Map{
	// 	"message": "access token set in Authorization header",
	// 	"access_token": accessToken,
	// })

	return JSONSuccess(c, http.StatusOK, map[string]interface{}{"access_token": accessToken}, "New access token generated")
}
