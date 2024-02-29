package services

import (
	"log"
	"net/http"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/ESMO-ENTERPRISE/auth-server/dtos"
	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Conn *database.Connector
}

func (us *UserService) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Query("email")

	var user models.User
	err := us.Conn.DB.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}

func (us *UserService) CreateUser(c *fiber.Ctx) error {
	payload := new(models.User)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Bad request")
	}

	var count int
	err := us.Conn.DB.QueryRow("SELECT COUNT(email) FROM users WHERE email = $1", payload.Email).Scan(&count)
	if 0 != 0 {
		return c.JSON(http.StatusBadRequest, "user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	result, err := us.Conn.DB.Exec("INSERT INTO users(name, email, password) VALUES($1, $2, $3)", payload.Name, payload.Email, hashedPassword)
	if err != nil {
		log.Fatal(err.Error())
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	id, _ := result.LastInsertId()
	var user models.User
	err = us.Conn.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user)
	if err != nil {
		log.Fatal(err.Error())
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(&user)
}

func (us *UserService) AssignateUserToClient(c *fiber.Ctx) error {
	payload := new(dtos.UserClientRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	}

	var user models.User
	if err := us.Conn.DB.QueryRow("SELECT * FROM users WHERE email = $1", payload.Email).Scan(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	}

	var client models.Client
	if err := us.Conn.DB.QueryRow("SELECT * FROM clients WHERE client_id = $1", payload.ClientID).Scan(&client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	}

	var count int
	if err := us.Conn.DB.QueryRow("SELECT COUNT(user_id) FROM users_clients WHERE user_id = $1 AND client_id = $2", user.ID, client.ID).Scan(&count); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	// User is assigned to client, no need to do something
	if count != 0 {
		return c.SendStatus(fiber.StatusOK)
	}

	if _, err := us.Conn.DB.Exec("INSERT INTO users_client(user_id, client_id) VALUES($1, $2)",
		user.ID, client.ID); err != nil {
		log.Fatal(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	rows, err := us.Conn.DB.Query("SELECT c.* FROM users_clients uc LEFT JOIN clients c ON uc.client_id = c.id WHERE user_id = $1 AND client_id = $2", user.ID, client.ID)
	if err != nil {
		log.Fatal(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		err := rows.Scan(&client)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, client)
	}

	return c.SendStatus(fiber.StatusOK)
}
