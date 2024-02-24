package routes

import (
	"github.com/ESMO-ENTERPRISE/auth-server/providers"
	"github.com/gofiber/fiber/v2"
)

func GithubRoutes(app *fiber.App) {
	githubRoutes := app.Group("/github")
	githubRoutes.Get("/login", providers.GithubLoginHandler)
	githubRoutes.Get("/callback", providers.GithubCallbackHandler)
}
