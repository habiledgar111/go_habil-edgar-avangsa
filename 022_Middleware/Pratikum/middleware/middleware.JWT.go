package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo"
)

func ExtractJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"massage": "JWT token invalid",
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"massage": "failed cast JWT",
			})
		}
		return c.JSON(http.StatusOK, claims)
	}
}
