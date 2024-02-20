package dtos

type RegisterPayload struct {
	Name     string `json:"name" validate:"required,min=3,max=60"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=200"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=200"`
}
