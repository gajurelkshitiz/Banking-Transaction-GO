package middlewares

import (
	"banking_transaction_go/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
		}
		if strings.HasPrefix(strings.ToLower(auth), "bearer ") {
			auth = auth[7:]
		}
		claims, err := utils.ValidateToken(auth)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		// store userID in context
		c.Set("user_id", claims.UserID)
		return next(c)
	}
}
