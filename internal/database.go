package internal

import (
	"log"
	"os"

	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func InitDatabase() {
	databaseUrl := os.Getenv("DATABASE_URL")

	DB, err := gorm.Open(postgres.Open(databaseUrl))
	if err != nil {
		log.Fatalln(err)
	}

	// Migrate
	DB.AutoMigrate(&models.Client{}, &models.Scope{})

	// Seed
	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "redirect_uri", "logo", "client_id", "client_secret"}),
	}).Create(&models.Client{
		ID:           uuid.New(),
		Name:         "Test client",
		ClientID:     "test-client",
		ClientSecret: "test-client",
		RedirectURI:  "http://localhost:3000/github/callback",
		Logo:         "https://placehold.co/600x400",
	})
}
