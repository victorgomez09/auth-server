package dtos

import "github.com/golang-jwt/jwt"

type TokenDto struct {
	*jwt.StandardClaims
	TokenType string
	UserId    string
	Scopes    []string
}
