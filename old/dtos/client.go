package dtos

import "github.com/google/uuid"

type ClientDto struct {
	ID           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	Logo         string     `json:"logo"`
	ClientID     string     `json:"clientID"`
	ClientSecret string     `json:"clientSecret"`
	RedirectURI  string     `json:"redirectURI"`
	Scopes       []ScopeDto `json:"scopes"`
}

type ScopeDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ClientRefer uuid.UUID `json:"clientID"`
}
