package utils

// import (
// 	"os"
// 	"strconv"

// 	"github.com/ESMO-ENTERPRISE/auth-server/database"
// 	"github.com/ESMO-ENTERPRISE/auth-server/models"
// 	"github.com/ESMO-ENTERPRISE/auth-server/token"
// 	"github.com/gofiber/fiber/v2"
// )

// func GenerateCookie(user *models.User, c *fiber.Ctx, conn *database.Connector) error {
// 	// refreshTokenUtil := token.RefreshToken{
// 	// 	Conn:   conn,
// 	// 	UserID: user.ID,
// 	// }
// 	accessTokenUtil := token.AccessToken{
// 		Conn:   conn,
// 		UserID: user.UserID,
// 	}
// 	sessionTokenUtil := token.SessionToken{
// 		Conn: conn,
// 	}

// 	// refreshToken, errRefreshToken := refreshTokenUtil.Create()
// 	// if errRefreshToken != nil {
// 	// 	return errRefreshToken
// 	// }

// 	accessToken, errAccessToken := accessTokenUtil.Create()
// 	if errAccessToken != nil {
// 		return errAccessToken
// 	}

// 	sessionToken, errSessionToken := sessionTokenUtil.Create(*user)
// 	if errSessionToken != nil {
// 		return errSessionToken
// 	}

// 	maxAge, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_TIME"))
// 	// c.Cookie(&fiber.Cookie{
// 	// 	Name:     "refresh_token",
// 	// 	Value:    *refreshToken.Token,
// 	// 	Path:     "/",
// 	// 	MaxAge:   maxAge * 60,
// 	// 	Secure:   true,
// 	// 	HTTPOnly: false,
// 	// 	Domain:   "/",
// 	// })

// 	c.Cookie(&fiber.Cookie{
// 		Name:     "access_token",
// 		Value:    accessToken,
// 		Path:     "/",
// 		MaxAge:   maxAge * 60,
// 		Secure:   true,
// 		HTTPOnly: false,
// 		Domain:   "/",
// 	})

// 	c.Cookie(&fiber.Cookie{
// 		Name:     "session",
// 		Value:    *sessionToken.Token,
// 		Path:     "/",
// 		MaxAge:   maxAge * 60,
// 		Secure:   false,
// 		HTTPOnly: false,
// 		Domain:   "/",
// 	})

// 	return nil
// }
