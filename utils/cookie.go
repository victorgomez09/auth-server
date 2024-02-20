package utils

import (
	"net/http"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/ESMO-ENTERPRISE/auth-server/token"
	"github.com/labstack/echo/v4"
)

func GenerateCookie(user *models.User, c echo.Context, conn *database.Connector) error {
	refreshTokenUtil := token.RefreshToken{
		Conn:   conn,
		UserID: user.ID,
	}
	accessTokenUtil := token.AccessToken{
		Conn:   conn,
		UserID: user.ID,
	}
	sessionTokenUtil := token.SessionToken{
		Conn: conn,
	}

	refreshToken, errRefreshToken := refreshTokenUtil.Create()
	if errRefreshToken != nil {
		return errRefreshToken
	}

	accessToken, errAccessToken := accessTokenUtil.Create()
	if errAccessToken != nil {
		return errAccessToken
	}

	sessionToken, errSessionToken := sessionTokenUtil.Create(*user)
	if errSessionToken != nil {
		return errSessionToken
	}

	c.Cookie(&http.Cookie{
		Name:  "access_token",
		Value: *accessToken.Token,
		Path:  "/",
		// MaxAge: strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_TIME")) * 60,
		Secure:   true,
		HttpOnly: false,
		// Name:     "access_token",
		// Value:    *accessTokenD.Token,
		// Path:     "/",
		// MaxAge:   env.AccessTokenMaxAge * 60,
		// Secure:   false,
		// HTTPOnly: false,
		// Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    *refreshTokenD.Token,
		Path:     "/",
		MaxAge:   env.RefreshTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    *sessionTokenD.Token,
		Path:     "/",
		MaxAge:   env.RefreshTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

	return nil
}
