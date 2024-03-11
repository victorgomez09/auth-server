package models

type User struct {
	ID       string    `json:"userID"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password []byte    `json:"password"`
	Clients  []*Client `json:"clients"`
}
