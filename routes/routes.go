package routes

import "github.com/labstack/echo/v4"

func InitRoutes(app *echo.Echo) {
	GithubRoutes(app)
}