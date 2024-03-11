package services

import (
	"log"
	"net/http"
	"os"
	"slices"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/ESMO-ENTERPRISE/auth-server/dtos"
	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/ESMO-ENTERPRISE/auth-server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Conn *database.Connector
	Sess *session.Store
}

func (as *AuthService) HandleAuthorize(c *fiber.Ctx) error {
	authContext := dtos.AuthContext{
		ClientId:            c.Query("client_id"),
		RedirectURI:         c.Query("redirect_uri"),
		ResponseType:        c.Query("response_type"),
		CodeChallengeMethod: c.Query("code_challenge_method"),
		CodeChallenge:       c.Query("code_challenge"),
		ResponseMode:        c.Query("response_mode"),
		MaxAge:              c.Query("max_age"),
		RequestedAcrValues:  c.Query("acr_values"),
		State:               c.Query("state"),
		Nonce:               c.Query("nonce"),
		// UserAgent:           c.UserAgent(),
		// IpAddress:           c.RemoteAddr,
	}

	// Validate client
	var client models.Client
	err := as.Conn.DB.QueryRow("SELECT * FROM clients WHERE client_id = $1", authContext.ClientId).Scan(&client)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if !client.Enabled {
		return c.Status(fiber.StatusBadRequest).JSON("The client associated with the provided client_id is not enabled.")
	}

	// payload redirect url match with client redirect url
	if authContext.RedirectURI != client.RedirectURI {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid redirect_uri parameter. The client does not have this redirect uri configured.")
	}

	// Validate response type
	if authContext.ResponseType != "code" {
		return c.Status(fiber.StatusBadRequest).JSON("Ensure response_type is set to 'code' as it's the only supported value.!")
	}

	// Validate challenge method
	if authContext.CodeChallengeMethod != "S256" {
		return c.Status(fiber.StatusBadRequest).JSON("invalid_request", "Ensure code_challenge_method is set to 'S256' as it's the only supported value.")
	}

	// Validate code challenge
	if len(authContext.CodeChallenge) < 43 || len(authContext.CodeChallenge) > 128 {
		return c.Status(fiber.StatusBadRequest).JSON("invalid_request", "The code_challenge parameter is either missing or incorrect. It should be 43 to 128 characters long.")
	}

	// Validate response mode
	if len(authContext.ResponseMode) > 0 {
		if !slices.Contains([]string{"query", "fragment", "form_post"}, authContext.ResponseMode) {
			return c.Status(fiber.StatusBadRequest).JSON("invalid_request", "Please use 'query,' 'fragment,' or 'form_post' as the response_mode value.")
		}
	}

	// Validate scopes
	validator := &utils.Oauth2Validator{
		Conn: as.Conn,
	}
	err = validator.ValidateScopes(c.Context(), authContext.Scope)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// TODO: check if user is in session to redirect

	// else, create new session
	sessionService := &utils.SessionService{
		Sess: as.Sess,
	}
	err = sessionService.SetContextSession(c, &authContext)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Cannot set session")
	}

	// Generate code
	v, _ := utils.CreateCodeVerifier()
	v.Value = authContext.

	frontendUrl := os.Getenv("FRONTEND_URL")
	return c.Status(fiber.StatusTemporaryRedirect).Redirect(frontendUrl + "/login")
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
