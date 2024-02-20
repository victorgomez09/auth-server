package token

import (
	"encoding/base64"
	"os"
	"time"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// Details is a struct that contains the data that need to be used when creating tokens
type Details struct {
	Token     *string
	ExpiresIn *int64
	TokenUUID string
	UserID    string
}

// RefreshToken is a struct that is used to perform operations on refresh tokens
type RefreshToken struct {
	Conn   *database.Connector
	UserID uuid.UUID
}

// Create a refresh token
func (r *RefreshToken) Create() (tokenDetails *Details, err error) {
	now := time.Now().UTC()

	tokenUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	tokenDetails = &Details{
		ExpiresIn: new(int64),
		Token:     new(string),
	}

	duration, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRATION_TIME"))
	if err != nil {
		return nil, err
	}
	*tokenDetails.ExpiresIn = now.Add(duration).Unix()
	tokenDetails.TokenUUID = tokenUUID.String()
	tokenDetails.UserID = r.UserID.String()

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(os.Getenv("REFRESH_TOKEN_PRIVATE_KEY"))
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return nil, err
	}

	claims := make(jwt.MapClaims)
	claims["sub"] = r.UserID.String()
	claims["token_uuid"] = tokenDetails.TokenUUID
	claims["exp"] = *tokenDetails.ExpiresIn
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	*tokenDetails.Token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return nil, err
	}

	// tokenVal, err := json.Marshal(schemas.RefreshTokenDetails{
	// 	UserID:          r.UserID.String(),
	// 	AccessTokenUUID: "",
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// err = r.Conn.R.Session.Set(context.TODO(), tokenDetails.TokenUUID, string(tokenVal), time.Unix(*tokenDetails.ExpiresIn, 0).Sub(now)).Err()
	return tokenDetails, err
}
