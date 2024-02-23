package services

import (
	"fmt"
	"net/http"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/ESMO-ENTERPRISE/auth-server/dtos"
	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/ESMO-ENTERPRISE/auth-server/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Conn *database.Connector
}

func (a *Auth) RegisterWithEmailAndPassword(c *fiber.Ctx) error {
	u := new(dtos.RegisterPayload)

	if err := c.BodyParser(&u); err != nil {
		fmt.Errorf(err.Error())

		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Errorf(err.Error())
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	user := &models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: hashedPassword,
	}

	if userExist := a.Conn.DB.Where("email = ?", user.Email).First(user).RowsAffected; userExist != 0 {
		return c.JSON(http.StatusBadRequest, "user already exists")
	}

	a.Conn.DB.Create(&user)

	return c.JSON(map[string]string{
		"data": "user " + user.Name + " created",
	})
}

func (a *Auth) LoginWithEmailAndPassword(c *fiber.Ctx) error {
	u := new(dtos.LoginPayload)

	if err := c.BodyParser(&u); err != nil {
		fmt.Errorf(err.Error())

		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	}

	user := &models.User{
		Email: u.Email,
	}

	err := a.Conn.DB.Where("email = ?", u.Email).First(user).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("bad credentials")
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if errCompare != nil {
		return c.Status(http.StatusBadRequest).JSON("bad credentials")
	}

	cookieError := utils.GenerateCookie(user, c, a.Conn)
	if cookieError != nil {
		return c.Status(http.StatusInternalServerError).JSON("error generating cookies")
	}

	utils.SetSession(user, c)

	return c.Status(fiber.StatusOK).JSON(user)
}

// func (a *Auth) Logout(c *fiber.Ctx) error {
// 	refreshToken := utils.GetRefreshToken(c)

// 	user := utils.Get(c)

// 	tokenService := token.RefreshToken{
// 		Conn:   a.Conn,
// 		UserID: user.ID,
// 	}

// 	refreshToken, err := tokenService.Create()

// }
