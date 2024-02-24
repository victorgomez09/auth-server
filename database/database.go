package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connector struct {
	DB *gorm.DB
}

func (c *Connector) InitDatabase() {
	databseUrl := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(databseUrl))
	if err != nil {
		log.Fatal(err)
	}

	c.DB = db
}

func (c *Connector) MigrateDatabase() {

	c.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	err := c.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print("\n\n ✅ All schema changes have been migrated !")
}
