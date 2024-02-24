package services

import (
	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/gofiber/fiber/v2"
)

type Conn struct {
	Conn *database.Connector
}

func (conn *Conn) Create(c *fiber.Ctx) error {
	return nil
}
