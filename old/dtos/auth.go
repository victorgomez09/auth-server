package dtos

import "time"

type AuthContext struct {
	ClientId            string
	RedirectURI         string
	ResponseType        string
	CodeChallengeMethod string
	CodeChallenge       string
	ResponseMode        string
	Scope               string
	ConsentedScope      string
	MaxAge              string
	RequestedAcrValues  string
	State               string
	Nonce               string
	UserAgent           string
	IpAddress           string
	AcrLevel            string
	AuthMethods         string
	AuthTime            time.Time
	UserId              int64
	AuthCompleted       bool
}

type RegisterPayload struct {
	Name     string `json:"name" validate:"required,min=3,max=60"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=200"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=200"`
	ClientID string `json:"clientID" validate:"required"`
}
