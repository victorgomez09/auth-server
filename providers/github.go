package providers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ESMO-ENTERPRISE/auth-server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var GithubAdress = "http://localhost:3000"
var GithubCallbackEndpoint = "/github/callback/"
var lf *loginFlow

type loginFlow struct {
	config *oauth2.Config
}

func InitGithubFlow() {
	githubClientID := os.Getenv("GITHUB_CLIENT_ID")
	githubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	if len(githubClientID) == 0 || len(githubClientSecret) == 0 {
		log.Fatal("Set GITHUB_CLIENT_* env vars")
	}

	config := &oauth2.Config{
		ClientID:     githubClientID,
		ClientSecret: githubClientSecret,
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
	}

	lf = &loginFlow{
		config: config,
	}
}

func GithubLoginHandler(c *fiber.Ctx) error {
	// Generate a random state for CSRF protection and set it in a cookie.
	state, err := utils.RandString(16)
	if err != nil {
		panic(err)
	}

	cookie := &fiber.Cookie{
		Name:     "state",
		Value:    state,
		Path:     "/",
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   c.Protocol() == "https",
		HTTPOnly: true,
	}
	fmt.Print(*cookie)
	c.Cookie(cookie)

	redirectUrl := lf.config.AuthCodeURL(state)
	c.Redirect(redirectUrl, fiber.StatusTemporaryRedirect)

	return nil
}

func GithubCallbackHandler(c *fiber.Ctx) error {
	state := c.Cookies("state")
	if state == "" {
		c.JSON(http.StatusBadRequest, "state not found")
	}

	if c.Query("state") != state {
		c.JSON(http.StatusBadRequest, "state values did not match")
	}

	code := c.Query("code")
	tok, err := lf.config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatal(err)
	}

	// This client will have a bearer token to access the GitHub API on
	// the user's behalf.
	client := lf.config.Client(context.Background(), tok)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		panic(err)
	}
	respbody, _ := io.ReadAll(resp.Body)
	userInfo := string(respbody)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userInfo,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		log.Fatal(err)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   c.Protocol() == "https",
		HTTPOnly: true,
	})
	c.Response().Header.Add("Content-type", "application/json")
	c.Redirect("http://localhost:5173", fiber.StatusPermanentRedirect)

	return nil
}
