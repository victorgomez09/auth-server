package main

import (
	"log"
	"net/http"

	"github.com/ESMO-ENTERPRISE/auth-server/providers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	providers.InitGithubFlow()

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	githubRoutes := app.Group("/github")
	githubRoutes.GET("", providers.GithubLoginHandler)
	githubRoutes.GET("/callback", providers.GithubCallbackHandler)

	log.Fatal(app.Start(":3000"))
}
