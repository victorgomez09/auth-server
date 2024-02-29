package services

import (
	"log"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/ESMO-ENTERPRISE/auth-server/dtos"
	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/ESMO-ENTERPRISE/auth-server/utils"
	"github.com/gofiber/fiber/v2"
)

type ClientService struct {
	Conn *database.Connector
}

func (cs *ClientService) CreateClient(c *fiber.Ctx) error {
	payload := new(dtos.ClientDto)
	if err := c.BodyParser(payload); err != nil {
		log.Fatal(err)

		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	}
	// var scopes []models.Scope
	// if len(payload.Scopes) > 0 {
	// 	for i := 0; i < len(payload.Scopes); i++ {
	// 		scopePayload := new(dtos.ScopeDto)
	// 		if err := c.BodyParser(payload.Scopes[i]); err != nil {
	// 			log.Fatal(err)

	// 			return c.Status(fiber.StatusBadRequest).JSON("bad request")
	// 		}

	// 		scopes = append(scopes, models.Scope{
	// 			Name:        scopePayload.Name,
	// 			ClientRefer: client.ID,
	// 		})
	// 	}
	// } else {
	// 	scopes = []models.Scope{}
	// }

	// client.Scopes = scopes

	var count int64
	cs.Conn.DB.QueryRow("SElECT COUNT(name) FROM clients WHERE name = $1", payload.Name).Scan(&count)
	if count != 0 {
		return c.Status(fiber.StatusBadRequest).JSON("client already exists")
	}

	// client.ClientID = utils.RandStringRunes(10)
	// client.ClientSecret = utils.RandStringRunes(20)
	// cli.Conn.DB.Create(&client)

	result, err := cs.Conn.DB.Exec("INSERT INTO clients(name, logo, client_id, client_secret, redirect_url) VALUES($1, $2, $3, $4, $5)",
		payload.Name, payload.Logo, utils.RandStringRunes(10), utils.RandStringRunes(20), payload.RedirectURI)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("client already exists")
	}

	id, _ := result.LastInsertId()
	var client models.Client
	err = cs.Conn.DB.QueryRow("SELECT * FROM clients WHERE id = $1", id).Scan(&client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	return c.Status(fiber.StatusCreated).JSON(&client)
}
