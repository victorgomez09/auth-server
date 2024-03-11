package utils

import (
	"encoding/json"

	"github.com/ESMO-ENTERPRISE/auth-server/dtos"
	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type SessionService struct {
	Sess *session.Store
}

func (ss *SessionService) SetContextSession(c *fiber.Ctx, authContext *dtos.AuthContext) error {
	sess, err := ss.Sess.Get(c)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(authContext)
	if err != nil {
		return err
	}

	sess.Set(SessionKeyAuthContext, string(jsonData))
	err = sess.Save()
	if err != nil {
		return err
	}

	return nil
}

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
