package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDatabase() (*sql.DB, error) {
	db, err := initDatabase()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initDatabase() (*sql.DB, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	// if err != nil {
	// 	return nil, err
	// }
	// db := client.Database(os.Getenv("MONGO_DB"))

	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
	return db, nil
}
