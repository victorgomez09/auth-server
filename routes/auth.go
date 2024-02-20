package routes

import (
	"github.com/ESMO-ENTERPRISE/auth-server/services"
	"github.com/labstack/echo/v4"
)

var service *services.Auth

func AuthRoutes(app *echo.Echo) {

	authRoutes := app.Group("/auth")
	authRoutes.POST("/register", service.RegisterWithEmailAndPassword)
	authRoutes.POST("/login")
	authRoutes.POST("/logout")
	authRoutes.POST("/refresh")
	reAuthRoutes := authRoutes.Group("/reAuthenticate")
	reAuthRoutes.POST("/password")
	reAuthRoutes.POST("/passkey")
}
