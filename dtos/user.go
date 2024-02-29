package dtos

import "github.com/google/uuid"

type UserResponse struct {
	ID      uuid.UUID   `json:"id"`
	Name    string      `json:"name"`
	Email   string      `json:"email"`
	Clients []ClientDto `json:"clients"`
}

type UserClientRequest struct {
	Email    string `json:"email"`
	ClientID string `json:"clientID"`
}

type UserClientResponse struct {
	User    UserResponse `json:"user"`
	Clients []ClientDto  `json:"clients"`
}
