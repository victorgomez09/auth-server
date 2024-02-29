package services

import (
	"log"
	"net/http"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/ESMO-ENTERPRISE/auth-server/dtos"
	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/ESMO-ENTERPRISE/auth-server/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Conn *database.Connector
}

func (as *AuthService) RegisterWithEmailAndPassword(c *fiber.Ctx) error {
	u := new(dtos.RegisterPayload)

	if err := c.BodyParser(&u); err != nil {
		log.Fatal(err.Error())

		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	}

	var cnt int
	as.Conn.DB.QueryRow("SELECT COUNT(email) FROM users WHERE email = $1", u.Email).Scan(&cnt)
	if cnt != 0 {
		return c.JSON(http.StatusBadRequest, "user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}
	_, err = as.Conn.DB.Exec("INSERT INTO users(name, email, password) VALUES($1, $2, $3)",
		u.Name, u.Email, hashedPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	return c.Status(fiber.StatusCreated).JSON("user created")
}

func (as *AuthService) LoginWithEmailAndPassword(c *fiber.Ctx) error {
	u := new(dtos.LoginPayload)

	if err := c.BodyParser(&u); err != nil {
		log.Fatal(err.Error())

		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	}

	// TODO: check if user is in client to continue
	// if count := a.Conn.DB.Where("client_id = ?", client.ClientID).First(client).RowsAffected; count == 0 {
	// 	return c.Status(fiber.StatusForbidden).JSON("user dont have permissions")
	// }

	var user models.User
	err := as.Conn.DB.QueryRow("SELECT * FROM users WHERE email = $1").Scan(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("bad credentials")
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if errCompare != nil {
		return c.Status(http.StatusBadRequest).JSON("bad credentials")
	}

	cookieError := utils.GenerateCookie(&user, c, as.Conn)
	if cookieError != nil {
		return c.Status(http.StatusInternalServerError).JSON("error generating cookies")
	}

	utils.SetSession(&user, c)

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
