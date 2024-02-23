package token

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/ESMO-ENTERPRISE/auth-server/database"
// 	"github.com/golang-jwt/jwt"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// // Details is a struct that contains the data that need to be used when creating tokens
// type Details struct {
// 	Token     *string
// 	ExpiresIn *int64
// 	TokenUUID string
// 	UserID    string
// }

// // RefreshToken is a struct that is used to perform operations on refresh tokens
// type RefreshToken struct {
// 	Conn   *database.Connector
// 	UserID uuid.UUID
// }

// // Get is a function that is used to get the refesh token details while verifying it
// func (r *RefreshToken) Get(tokenStr string) (token *Details, err error) {
// 	tokenDetails, _, err := validate(r.Conn, tokenStr, os.Getenv("REFRESH_TOKEN_PUBLIC_KEY"), r.UserID.String())
// 	if err != nil {
// 		if err == errors.ErrUnauthorized {
// 			return nil, errors.ErrRefreshTokenExpired
// 		}
// 		return nil, err
// 	}

// 	tokenUUID, err := uuid.Parse(tokenDetails.TokenUUID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var session models.Sessions
// 	err = r.Conn.DB.Where(&models.Sessions{
// 		ID: &tokenUUID,
// 	}).First(&session).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, errors.ErrRefreshTokenExpired
// 		}

// 		return nil, err
// 	}

// 	return tokenDetails, nil
// }

// // Create a refresh token
// func (r *RefreshToken) Create() (tokenDetails *Details, err error) {
// 	now := time.Now().UTC()

// 	tokenUUID, err := uuid.NewUUID()
// 	if err != nil {
// 		return nil, err
// 	}

// 	tokenDetails = &Details{
// 		ExpiresIn: new(int64),
// 		Token:     new(string),
// 	}

// 	duration, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRATION_TIME"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	*tokenDetails.ExpiresIn = now.Add(duration).Unix()
// 	tokenDetails.TokenUUID = tokenUUID.String()
// 	tokenDetails.UserID = r.UserID.String()

// 	decodedPrivateKey, err := base64.StdEncoding.DecodeString(os.Getenv("REFRESH_TOKEN_PRIVATE_KEY"))
// 	if err != nil {
// 		return nil, err
// 	}

// 	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	claims := make(jwt.MapClaims)
// 	claims["sub"] = r.UserID.String()
// 	claims["token_uuid"] = tokenDetails.TokenUUID
// 	claims["exp"] = *tokenDetails.ExpiresIn
// 	claims["iat"] = now.Unix()
// 	claims["nbf"] = now.Unix()

// 	*tokenDetails.Token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// tokenVal, err := json.Marshal(schemas.RefreshTokenDetails{
// 	// 	UserID:          r.UserID.String(),
// 	// 	AccessTokenUUID: "",
// 	// })
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// err = r.Conn.R.Session.Set(context.TODO(), tokenDetails.TokenUUID, string(tokenVal), time.Unix(*tokenDetails.ExpiresIn, 0).Sub(now)).Err()
// 	return tokenDetails, err
// }

// func validate(conn *database.Connector, token, publicKey, userID string) (tokenDetails *Details, metadata interface{}, err error) {
// 	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, fmt.Errorf("unexpected method : %s", t.Header["alg"])
// 		}

// 		return key, nil
// 	})
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	claims, ok := parsedToken.Claims.(jwt.MapClaims)
// 	if !ok || !parsedToken.Valid {
// 		return nil, nil, fmt.Errorf("validate : invalid token")
// 	}

// 	exp := int64(claims["exp"].(float64))
// 	tokenDetails = &Details{
// 		TokenUUID: fmt.Sprint(claims["token_uuid"]),
// 		UserID:    fmt.Sprint(claims["sub"]),
// 		ExpiresIn: &exp,
// 		Token:     &token,
// 	}
// 	if tokenDetails.UserID != userID {
// 		return nil, nil, err
// 	}

// 	// val := conn.R.Session.Get(context.TODO(), tokenDetails.TokenUUID).Val()
// 	// if val == "" {
// 	// 	return nil, nil, errors.ErrUnauthorized
// 	// }

// 	// var valStr map[string]interface{}
// 	// err = json.Unmarshal([]byte(val), &valStr)
// 	// if err != nil {
// 	// 	return tokenDetails, nil, nil
// 	// }

// 	// metadata = schemas.RefreshTokenDetails{
// 	// 	UserID:          valStr["UserID"].(string),
// 	// 	AccessTokenUUID: valStr["AccessTokenUUID"].(string),
// 	// }

// 	// now := time.Now().UTC().Unix()
// 	// if *tokenDetails.ExpiresIn <= now {
// 	// 	return nil, nil, errors.ErrUnauthorized
// 	// }

// 	return tokenDetails, metadata, nil
// }
