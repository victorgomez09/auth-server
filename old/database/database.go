package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Connector struct {
	DB *sql.DB
}

func (c *Connector) InitDatabase() {
	// databseUrl := os.Getenv("DATABASE_URL")
	// db, err := gorm.Open(postgres.Open(databseUrl))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// c.DB = db

	connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	c.DB = db
}

// func (c *Connector) MigrateDatabase() {

// 	c.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

// 	err := c.DB.AutoMigrate(&models.User{}, &models.Client{})
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	fmt.Print("\n\n ✅ All schema changes have been migrated !")
// }
