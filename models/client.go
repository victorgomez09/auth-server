package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primaryKey,type:uuid;default:uuid_generate_v4()"`
	Name         string
	Logo         string
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scopes       []Scope
}

type Scope struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey,type:uuid;default:uuid_generate_v4()"`
	Name        string
	ClientRefer uuid.UUID
}
