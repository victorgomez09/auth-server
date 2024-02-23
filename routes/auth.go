package routes

import (
	"github.com/ESMO-ENTERPRISE/auth-server/services"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(service *services.Auth, app *fiber.App) {

	authRoutes := app.Group("/auth")
	authRoutes.Post("/register", service.RegisterWithEmailAndPassword)
	authRoutes.Post("/login", service.LoginWithEmailAndPassword)
	// authRoutes.Post("/logout")
	// authRoutes.Post("/refresh")

	// reAuthRoutes := authRoutes.Group("/reAuthenticate")
	// reAuthRoutes.Post("/password")
	// reAuthRoutes.Post("/passkey")
}
