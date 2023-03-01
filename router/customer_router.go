package router

import (
	"domain-driven-design/config"
	"domain-driven-design/infrastructure/transport/http_transport"
	"domain-driven-design/middleware"
	"github.com/labstack/echo/v4"
)

func InitCustomerRouter(g *echo.Group, handler *http_transport.CustomerHandler, config *config.Config) {
	g.Use(middleware.JWTAuth(config))
	g.GET("", getAll(handler))
	g.GET("/:id", getOne(handler))
	g.POST("", create(handler))
}

func getAll(handler *http_transport.CustomerHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler.GetCustomers(c)
	}
}

func getOne(handler *http_transport.CustomerHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler.GetCustomer(c)
	}
}

func create(handler *http_transport.CustomerHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler.CreateCustomer(c)
	}
}
