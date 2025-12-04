package controllers

import "github.com/labstack/echo/v4"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func JSONSuccess(c echo.Context, status int, data interface{}, message string) error {
	return c.JSON(status, Response{Code: status, Message: message, Data: data})
}

func JSONError(c echo.Context, status int, errMsg, message string) error {
	return c.JSON(status, Response{Code: status, Message: message, Error: errMsg})
}
