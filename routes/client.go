package routes

import "github.com/gofiber/fiber"

func ClientRoutes(app *fiber.App) {
	clientRoutes := app.Group("/client")

	clientRoutes.Post("/register", services.Client.Register)
	clientRoutes.Post("/login", services.Client.Login)
}
