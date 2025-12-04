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
	var body RegisterRequest
	if err := c.Bind(&body); err != nil {
		return JSONError(c, http.StatusBadRequest, "Invalid request", "")
	}
	if body.Name == "" || body.Email == "" || body.Password == "" {
		return JSONError(c, http.StatusBadRequest, "missing fields", "")
	}

	user, err := ac.Service.Register(body.Name, body.Email, body.Password)
	if err != nil {
		return JSONError(c, http.StatusBadRequest, err.Error(), "")
	}

	return JSONSuccess(c, http.StatusCreated, map[string]interface{}{"user": user}, "User registered successfully")
}

func (ac *AuthController) Login(c echo.Context) error {
	var body LoginRequest
	if err := c.Bind(&body); err != nil {
		return JSONError(c, http.StatusBadRequest, "Invalid request", "")
	}
	if body.Email == "" || body.Password == "" {
		return JSONError(c, http.StatusBadRequest, "missing credentials", "")
	}

	access, refresh, user, err := ac.Service.Login(body.Email, body.Password)
	if err != nil {
		return JSONError(c, http.StatusUnauthorized, err.Error(), "")
	}

	return JSONSuccess(c, http.StatusOK, AuthPayload{
		AccessToken:  access,
		RefreshToken: refresh,
		User:         user,
	}, "")
}

func (ac *AuthController) Refresh(c echo.Context) error {
	var body RefreshRequest
	if err := c.Bind(&body); err != nil {
		return JSONError(c, http.StatusBadRequest, "Invalid request", "")
	}
	if body.RefreshToken == "" {
		return JSONError(c, http.StatusBadRequest, "missing refresh token", "")
	}

	accessToken, err := ac.Service.Refresh(body.RefreshToken)
	if err != nil {
		return JSONError(c, http.StatusUnauthorized, err.Error(), "")
	}

	return JSONSuccess(c, http.StatusOK, map[string]interface{}{"access_token": accessToken}, "New access token generated")
}
