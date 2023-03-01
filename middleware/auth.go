package middleware

import (
	"domain-driven-design/config"
	"domain-driven-design/pkg/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func JWTAuth(config *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("token")
			if token == "" {
				return c.String(http.StatusUnauthorized, "token not found")
			}

			_, err := jwt.VerifyToken(token, config.JwtSecretKey)
			if err != nil {
				return c.String(http.StatusUnauthorized, "token invalid")
			}

			return next(c)
		}
	}
}
