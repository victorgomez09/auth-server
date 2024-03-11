package model

import (
	"github.com/google/uuid"
	"github.com/miyuki-starmiya/go-oauth2-server/db/constants"
)

type AuthorizationData struct {
	ID                  uuid.UUID                      `db:"id"`
	ClientID            string                         `db:"client_id"`
	RedirectURI         string                         `db:"redirect_uri"`
	AuthorizationCode   string                         `db:"authorization_code"`
	CodeChallenge       *string                        `db:"code_challenge"`
	CodeChallengeMethod *constants.CodeChallengeMethod `db:"code_challenge_method"`
}

type AuthorizationDataOption func(*AuthorizationData)

func WithCodeChallenge(codeChallenge string) AuthorizationDataOption {
	return func(ad *AuthorizationData) {
		ad.CodeChallenge = &codeChallenge
	}
}

func WithCodeChallengeMethod(codeChallengeMethod constants.CodeChallengeMethod) AuthorizationDataOption {
	return func(ad *AuthorizationData) {
		ad.CodeChallengeMethod = &codeChallengeMethod
	}
}

func NewAuthorizationData(clientID, redirectURI, authorizationCode string, opts ...AuthorizationDataOption) *AuthorizationData {
	ad := &AuthorizationData{
		ClientID:          clientID,
		RedirectURI:       redirectURI,
		AuthorizationCode: authorizationCode,
	}
	for _, opt := range opts {
		opt(ad)
	}

	return ad
}
