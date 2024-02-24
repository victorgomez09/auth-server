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

// const (
// 	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa keysize
// 	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
// )

// var (
// 	verifyKey *rsa.PublicKey
// 	signKey   *rsa.PrivateKey
// )

var privateKey *rsa.PrivateKey

func InitTokenService() {
	// signBytes, err := ioutil.ReadFile(privKeyPath)
	// if err != nil {
	// 	fmt.Println("Error getting private key")
	// }

	// signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	// if err != nil {
	// 	fmt.Println("Error parsing privat  key")
	// }

	// verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	// if err != nil {
	// 	fmt.Println("Error getting public key")
	// }

	// verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	// if err != nil {
	// 	fmt.Println("Error verifying keys")
	// }

	// Just as a demo, generate a new private/public key pair on each run. See note above.
	rng := rand.Reader
	var err error
	privateKey, err = rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}
}

// Create is a function that is used to create the access token
func (a *AccessToken) Create() (string, error) {
	// t := jwt.New(jwt.GetSigningMethod(jwt.SigningMethodRS512.Name))

	// duration, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRATION_TIME"))
	// if err != nil {
	// 	fmt.Println("error parsing expiration time")
	// 	return "", err
	// }

	// // set our claims
	// t.Claims = &dtos.TokenDto{
	// 	&jwt.StandardClaims{
	// 		// set the expire time
	// 		// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
	// 		ExpiresAt: time.Now().Add(duration).Unix(),
	// 	},
	// 	"token_type",
	// 	a.UserID.String(),
	// 	[]string{"scope1", "scope2"},
	// }

	// return t.SignedString(signKey)

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
	t, err := token.SignedString(privateKey)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return "", err
	}

	return t, nil
}
