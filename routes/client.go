package routes

import (
	"github.com/ESMO-ENTERPRISE/auth-server/services"
	"github.com/gofiber/fiber/v2"
)

func ClientRoutes(service *services.Client, app *fiber.App) {
	clientRoutes := app.Group("/client")

	clientRoutes.Post("/create", service.CreateClient)
	// clientRoutes.Post("/login", services.Client.Login)
}
