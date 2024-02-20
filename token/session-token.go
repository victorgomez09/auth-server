package token

import (
	"os"
	"time"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// SessionToken is struct that manages the session token
type SessionToken struct {
	Conn *database.Connector
}

// Create is a function that is used to create a new session token
func (s *SessionToken) Create(user models.User) (tokenDetails *Details, err error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	tokenDetails = &Details{
		ExpiresIn: new(int64),
		Token:     new(string),
	}
	duration, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRATION_TIME"))
	if err != nil {
		return nil, err
	}
	*tokenDetails.ExpiresIn = now.Add(duration).Unix()
	tokenDetails.TokenUUID = uid.String()
	tokenDetails.UserID = user.ID.String()

	claims := make(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["token_uuid"] = tokenDetails.TokenUUID
	claims["exp"] = tokenDetails.ExpiresIn
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["name"] = user.Name
	claims["email"] = user.Email

	*tokenDetails.Token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SESSION_TOKEN_KEY")))
	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
}
