package utils

import (
	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/gofiber/fiber/v2"
)

func SetSession(user *models.User, c *fiber.Ctx) {
	c.Locals("user", user)
}

// Get the user details from the session
func Get(c *fiber.Ctx) (user *models.User) {
	return c.Locals("user").(*models.User)
}

func GetRefreshToken(c *fiber.Ctx) string {
	return c.Locals("refresh_token").(string)
}
