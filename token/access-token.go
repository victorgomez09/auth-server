package token

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// AccessToken is a struct that is used to perform operations on access tokens
type AccessToken struct {
	Conn   *database.Connector
	UserID uuid.UUID
}

// Create is a function that is used to create the access token
func (a *AccessToken) Create() (tokenDetails *Details, err error) {
	now := time.Now().UTC()

	tokenUUID, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("error generating uuid")
		return nil, err
	}

	tokenDetails = &Details{
		ExpiresIn: new(int64),
		Token:     new(string),
	}

	duration, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRATION_TIME"))
	if err != nil {
		fmt.Println("error parsing expiration time")
		return nil, err
	}

	*tokenDetails.ExpiresIn = now.Add(duration).Unix()
	tokenDetails.TokenUUID = tokenUUID.String()
	tokenDetails.UserID = a.UserID.String()

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"))
	if err != nil {
		fmt.Println("error decoding access token private key")
		return nil, err
	}

	fmt.Print(decodedPrivateKey)
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		fmt.Println("error parsing access token private key")
		fmt.Println(err)
		return nil, err
	}

	claims := make(jwt.MapClaims)
	claims["sub"] = a.UserID.String()
	claims["token_uuid"] = tokenDetails.TokenUUID
	claims["exp"] = *tokenDetails.ExpiresIn
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	*tokenDetails.Token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		fmt.Println("error generating token")
		return nil, err
	}

	// ctx := context.TODO()

	// detailsStr := a.Conn.R.Session.Get(ctx, refreshTokenUUID).Val()
	// if detailsStr != "" {
	// 	var details schemas.RefreshTokenDetails
	// 	err := json.Unmarshal([]byte(detailsStr), &details)
	// 	if err == nil {
	// 		a.Conn.R.Session.Del(ctx, details.AccessTokenUUID)
	// 	}
	// }

	// tokenVal, err := json.Marshal(schemas.RefreshTokenDetails{
	// 	UserID:          a.UserID.String(),
	// 	AccessTokenUUID: tokenDetails.TokenUUID,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// ttl := a.Conn.R.Session.TTL(ctx, refreshTokenUUID).Val()
	// if ttl.Seconds() < 0 {
	// 	ttl = 0
	// }
	// err = a.Conn.R.Session.Set(ctx, refreshTokenUUID, string(tokenVal), ttl).Err()
	// if err != nil {
	// 	return nil, err
	// }

	// err = a.Conn.R.Session.Set(ctx, tokenDetails.TokenUUID, a.UserID.String(), time.Unix(*tokenDetails.ExpiresIn, 0).Sub(now)).Err()
	return tokenDetails, err
}
