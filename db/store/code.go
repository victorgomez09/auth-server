package store

import (
	"context"
	"database/sql"
	"log"

	"github.com/miyuki-starmiya/go-oauth2-server/db/model"
)

func NewCodeStore(db *sql.DB) *CodeStore {
	return &CodeStore{
		DB: db,
	}
}

type CodeStore struct {
	DB *sql.DB
}

func (cs *CodeStore) CreateData(data *model.AuthorizationData) error {
	// collection := cs.DB.Collection("codes")

	// Insert a single document
	// insertResult, err := collection.InsertOne(context.TODO(), data)
	insertResult, err := cs.DB.ExecContext(context.TODO(), "INSERT INTO autorization_data(client_id, redirect_uri, authorization_code, code_challenge, code_challenge_method) VALUES($1, $2, $3, $4, $5)", data.ClientID, data.RedirectURI, data.CodeChallenge, data.CodeChallengeMethod)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return err
	}

	var code model.AuthorizationData
	id, _ := insertResult.LastInsertId()
	err2 := cs.DB.QueryRowContext(context.TODO(), "SELECT * from authorization_data WHERE id = $1", id).Scan(&code)
	if err2 != nil {
		log.Printf("Error: %v\n", err)
		return err
	}
	log.Printf("Inserted a single document: %v\n", code.ID)
	return nil
}

func (cs *CodeStore) GetData(clientId string, authorizationCode string) (*model.AuthorizationData, error) {
	// collection := cs.DB.Collection("codes")

	// // Define the filter criteria
	// filter := bson.M{
	// 	"client_id":          clientId,
	// 	"authorization_code": authorizationCode,
	// }

	var result *model.AuthorizationData
	err := cs.DB.QueryRowContext(context.TODO(), "SELECT * FROM authorization_data WHERE client_id = $1 AND authorization_code = $2", clientId, authorizationCode).Scan(&result)
	// err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}

	// Print the found document
	log.Printf("Found a document: %+v\n", result)
	return result, nil
}
