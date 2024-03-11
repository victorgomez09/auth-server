package model

import (
	"time"

	"github.com/google/uuid"
)

type TokenData struct {
	ID           uuid.UUID `db:"id"`
	ClientID     string    `db:"client_id"`
	AccessToken  string    `db:"access_token"`
	IssuedAt     time.Time `db:"issued_at"`
	ExpiresIn    int       `db:"expires_in"`
	RefreshToken string    `db:"refresh_token"`
	TokenType    string    `db:"token_type"`
}
