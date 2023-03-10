package http_transport

import (
	"context"
	"domain-driven-design/domain/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	LoginRes struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (ah *AuthHandler) Login(c echo.Context) error {
	var loginReq service.LoginReq
	err := c.Bind(&loginReq)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	ctx := context.Background()
	res, err := ah.authService.Login(ctx, &loginReq)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (ah *AuthHandler) RefreshAccessToken(c echo.Context) error {
	var refreshReq service.RefreshReq
	err := c.Bind(&refreshReq)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	//refreshToken := c.Request().Header.Get("refresh_token")
	ctx := context.Background()
	authInfo, err := ah.authService.RefreshAccessToken(ctx, &refreshReq)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, authInfo)
}
