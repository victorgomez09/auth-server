package utils

import (
	"sync"

	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/labstack/echo/v4"
)

var (
	Server *server.Server
	once   sync.Once
)

// InitServer Initialize the service
func InitServer(manager *manage.Manager) *server.Server {
	once.Do(func() {
		Server = server.NewDefaultServer(manager)
	})
	return Server
}

// HandleAuthorizeRequest the authorization request handling
func HandleAuthorizeRequest(c echo.Context) error {
	return Server.HandleAuthorizeRequest(c.Response().Writer, c.Request())
}

// HandleTokenRequest token request handling
func HandleTokenRequest(c echo.Context) error {
	return Server.HandleTokenRequest(c.Response().Writer, c.Request())
}
