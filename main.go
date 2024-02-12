package main

import (
	"log"
	"net/http"

	"github.com/ESMO-ENTERPRISE/auth-server/api"
	"github.com/ESMO-ENTERPRISE/auth-server/internal"
	"github.com/ESMO-ENTERPRISE/auth-server/providers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// call oauth2 init functions
	internal.InitOauth2()

	// Init github provider
	providers.InitGithubFlow()

	echoServer := echo.New()
	echoServer.HideBanner = true

	// Routes
	echoServer.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	echoServer.POST("/credentials", api.CreateCredentials)
	echoServer.GET("/token", api.GetToken)
	echoServer.GET("/protected", func(c echo.Context) error {
		return c.String(http.StatusOK, "/protected")
	}, internal.ValidateToken)

	echoServer.POST("/github-login", providers.GithubLoginHandler)
	echoServer.POST(providers.GithubCallbackEndpoint, providers.GithubCallbackHandler)

	echoServer.Logger.Fatal(echoServer.Start(":1209"))
}
