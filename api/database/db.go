package database

import (
	"again/api/models"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var DB = DbCon()

func DbCon() *sql.DB {
	constr := `host=localhost port=5432 user=postgres password=Pawan@2003 dbname=user sslmode=disable `
	DB, err := sql.Open("postgres", constr)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return DB
}

func CreateUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
        userID UUID PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL
     
    );`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	log.Println("Users table created successfully.")
}

func InsertUser(u models.User) (uuid.UUID, error) {
	query := `INSERT INTO users(userID, email, password) VALUES ($1, $2, $3)`
	userID := uuid.New()
	_, err := DB.Exec(query, userID, u.Email, u.Password)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert user: %v", err)
	}
	return userID, nil
}
