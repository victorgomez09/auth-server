package services

import (
	"log"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/ESMO-ENTERPRISE/auth-server/dtos"
	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/gofiber/fiber/v2"
)

type Client struct {
	Conn *database.Connector
}

func (cli *Client) CreateClient(c *fiber.Ctx) error {
	payload := new(dtos.ClientDto)
	if err := c.BodyParser(payload); err != nil {
		log.Fatal(err)

		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	}

	client := models.Client{
		Name:         payload.Name,
		Logo:         payload.Logo,
		ClientID:     payload.ClientID,
		ClientSecret: payload.ClientSecret,
		RedirectURI:  payload.RedirectURI,
	}

	var scopes []models.Scope
	if len(payload.Scopes) > 0 {
		for i := 0; i < len(payload.Scopes); i++ {
			scopePayload := new(dtos.ScopeDto)
			if err := c.BodyParser(payload.Scopes[i]); err != nil {
				log.Fatal(err)

				return c.Status(fiber.StatusBadRequest).JSON("bad request")
			}

			scopes = append(scopes, models.Scope{
				Name:        scopePayload.Name,
				ClientRefer: client.ID,
			})
		}
	} else {
		scopes = []models.Scope{}
	}

	client.Scopes = scopes

	var count int64
	if clientExists := cli.Conn.DB.Where("name = ?", client.Name).Count(&count); clientExists != nil {
		return c.Status(fiber.StatusBadRequest).JSON("client already exists")
	}

	cli.Conn.DB.Create(&client)

	return c.Status(fiber.StatusCreated).JSON(client)
}
