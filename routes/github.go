package routes

import (
	"github.com/ESMO-ENTERPRISE/auth-server/providers"
	"github.com/labstack/echo/v4"
)

func GithubRoutes(app *echo.Echo) {
	githubRoutes := app.Group("/github")
	githubRoutes.GET("/login", providers.GithubLoginHandler)
	githubRoutes.GET("/callback", providers.GithubCallbackHandler)
}