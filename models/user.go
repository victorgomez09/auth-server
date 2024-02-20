package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey,type:uuid;default:uuid_generate_v4()"`
	Name     string
	Email    string
	Password []byte
}
