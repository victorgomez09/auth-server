package store

import (
	"context"
	"database/sql"
	"log"

	"github.com/miyuki-starmiya/go-oauth2-server/db/model"
)

func NewTokenStore(db *sql.DB) *TokenStore {
	return &TokenStore{
		DB: db,
	}
}

type TokenStore struct {
	DB *sql.DB
}

func (ts *TokenStore) CreateData(data *model.TokenData) error {
	// collection := ts.DB.Collection("tokens")

	// Insert a single document
	// insertResult, err:= collection.InsertOne(context.TODO(), data)
	insertResult, err := ts.DB.ExecContext(context.TODO(), "INSERT INTO token_data(client_id, access_token, issued_at, expires_in, refresh_token, token_type) VALUES($1, $2, $3, $4, $5, $6)",
		data.ClientID, data.AccessToken, data.IssuedAt, data.ExpiresIn, data.RefreshToken, data.TokenType)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return err
	}

	id, _ := insertResult.LastInsertId()
	var token model.TokenData
	err2 := ts.DB.QueryRowContext(context.TODO(), "SELECT * FROM token_data WHERE id = $1", id).Scan(&token)
	if err2 != nil {
		log.Printf("Error: %v\n", err)
		return err
	}

	log.Printf("Inserted a single document: %v\n", token.ID)
	return nil
}

func (ts *TokenStore) GetData(clientId string, accessToken string) (*model.TokenData, error) {
	// collection := ts.DB.Collection("tokens")

	// // Define the filter criteria
	// filter := bson.M{
	// 	"client_id":    clientId,
	// 	"access_token": accessToken,
	// }

	var result *model.TokenData
	// err := collection.FindOne(context.TODO(), filter).Decode(&result)
	err := ts.DB.QueryRowContext(context.TODO(), "SELECT * FROM token_data WHERE client_id = $1 AND access_token = $2", clientId, accessToken).Scan(&result)
	if err != nil {
		// if err == mongo.ErrNoDocuments {
		// 	log.Println("No document was found with the given params")
		// } else {
		// 	log.Printf("Error: %v\n", err)
		// }
		log.Printf("Error: %v\n", err)
		return nil, err
	}

	// Print the found document
	log.Printf("Found a document: %+v\n", result)
	return result, nil
}
