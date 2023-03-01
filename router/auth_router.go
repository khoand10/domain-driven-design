package router

import (
	"domain-driven-design/infrastructure/transport/http_transport"
	"github.com/labstack/echo/v4"
)

func InitAuthRouter(g *echo.Group, handler *http_transport.AuthHandler) {
	g.POST("/login", login(handler))
	g.POST("/refresh", refreshAccessToken(handler))
}

func login(handler *http_transport.AuthHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler.Login(c)
	}
}

func refreshAccessToken(handler *http_transport.AuthHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler.RefreshAccessToken(c)
	}
}
