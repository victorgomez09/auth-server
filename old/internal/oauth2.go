package internal

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

var Oauth2Config *oauth2.Config

func InitOauth2Flow() {

	Oauth2Config = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"profile", "email"},
		RedirectURL:  "http://localhost:3000/callback",
	}
}

func Oauth2Login(c echo.Context) error {
	url := Oauth2Config.AuthCodeURL("state")

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func Oauth2Callback(c echo.Context) error {
	state := c.Request().FormValue("state")
	if state != "state" {
		c.JSON(http.StatusBadRequest, "Invalid state parameter")
	}

	code := c.Request().FormValue("code")
	token, err := Oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Failed to exchange token")
	}

	return c.JSON(http.StatusOK, token)
}
