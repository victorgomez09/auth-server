package api

import (
	"github.com/ESMO-ENTERPRISE/auth-server/internal"
	"github.com/labstack/echo/v4"
)

func GetToken(c echo.Context) error {
	return internal.Oauth2Server.HandleTokenRequest(c.Response().Writer, c.Request())
}
