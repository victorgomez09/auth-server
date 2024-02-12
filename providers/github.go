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
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var GithubAdress = "http://localhost:1209"
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

func GithubLoginHandler(c echo.Context) error {
	// Generate a random state for CSRF protection and set it in a cookie.
	state, err := utils.RandString(16)
	if err != nil {
		panic(err)
	}

	cookie := &http.Cookie{
		Name:     "state",
		Value:    state,
		Path:     "/",
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   c.Request().TLS != nil,
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	redirectUrl := lf.config.AuthCodeURL(state)
	return c.Redirect(http.StatusMovedPermanently, redirectUrl)
}

func GithubCallbackHandler(c echo.Context) error {
	state, err := c.Cookie("state")
	if err != nil {
		c.JSON(http.StatusBadRequest, "state not found")
	}

	if c.Request().URL.Query().Get("state") != state.Value {
		c.JSON(http.StatusBadRequest, "state values did not match")
	}

	code := c.Request().URL.Query().Get("code")
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

	c.Response().Header().Set("Content-type", "application/json")
	fmt.Println(string(userInfo))

	return c.JSON(http.StatusOK, tok.AccessToken)
}
