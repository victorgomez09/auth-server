package models

type Client struct {
	ID           string
	Name         string
	Logo         string
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Enabled      bool
	Users        []User
}
