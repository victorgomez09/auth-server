package api

import (
	"fmt"
	"net/http"

	"github.com/ESMO-ENTERPRISE/auth-server/internal"
	"github.com/labstack/echo/v4"
	"gopkg.in/oauth2.v3/models"
)

type Credential struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Domain       string `json:"domain"`
}

func CreateCredentials(c echo.Context) error {
	credential := new(Credential)

	if err := c.Bind(credential); err != nil {
		fmt.Println(err.Error())

		return err
	}

	err := internal.ClientStore.Set(credential.ClientId, &models.Client{
		ID:     credential.ClientId,
		Secret: credential.ClientSecret,
		Domain: credential.Domain,
	})

	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	return c.JSON(http.StatusCreated, credential)
}
