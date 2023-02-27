package router

import (
	"domain-driven-design/config"
	"domain-driven-design/infrastructure/transport/http_transport"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RestServer struct {
	Engine          *echo.Echo
	Config          *config.Config
	CustomerHandler *http_transport.CustomerHandler
}

func InitRouter(server *RestServer) error {
	apiVersion := getAPIVersion(server)
	group := server.Engine.Group(apiVersion)

	// Routes
	InitCustomerRouter(group.Group("/customers"), server.CustomerHandler)

	server.Engine.GET("", ping)

	port := getPort(server)
	err := server.Engine.Start(port)
	if err != nil {
		return err
	}
	return nil
}

func getAPIVersion(server *RestServer) string {
	return fmt.Sprintf("/%s", server.Config.APIVersion)
}

func getPort(server *RestServer) string {
	return fmt.Sprintf(":%s", server.Config.Port)
}

// TestHandler
func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
