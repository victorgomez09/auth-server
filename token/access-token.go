package token

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
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

var PrivateKey *rsa.PrivateKey

func InitTokenService() {
	rng := rand.Reader
	var err error
	PrivateKey, err = rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}
}

// Create is a function that is used to create the access token
func (a *AccessToken) Create() (string, error) {
	duration, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRATION_TIME"))
	if err != nil {
		fmt.Println("error parsing expiration time")
		return "", err
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":     a.UserID,
		"scopes": true,
		"exp":    time.Now().Add(duration * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(PrivateKey)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return "", err
	}

	return t, nil
}
