package routes

import "github.com/labstack/echo/v4"

func AuthRoutes(app *echo.Echo) {

	authRoutes := app.Group("/auth")
	authRoutes.GET("/authorize", )
}
