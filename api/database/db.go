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

func DeleteUser(id uuid.UUID) (*string, error) {
	// SQL query to delete a user by ID
	query := `DELETE FROM users WHERE userID = $1`

	// Execute the query
	result, err := DB.Exec(query, id)
	if err != nil {
		return nil, err
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		msg := "No user found with the provided ID"
		return &msg, nil
	}

	successMsg := "User deleted successfully"
	return &successMsg, nil
}

func Login(user models.User) (uuid.UUID, error) {
	query := `SELECT userID, password FROM users WHERE email = $1`

	var password string
	var id uuid.UUID

	// Execute the query
	err := DB.QueryRow(query, user.Email).Scan(&id, &password)
	if err != nil {
		if err == sql.ErrNoRows {
			return uuid.Nil, fmt.Errorf("user not found")
		}
		return uuid.Nil, err
	}

	// password verification logic here

	if user.Password != password {
		return uuid.Nil, fmt.Errorf("password mismatch")
	}

	return id, nil
}

func AllUsers() ([]models.User, error) {
    var users []models.User
    query := `SELECT userId, email, password FROM users`

    // Execute the query
    res, err := DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer res.Close()

    // Loop through the result set
    for res.Next() {
        var user models.User
        // Scan the result into the user struct
        err := res.Scan(&user.UserId, &user.Email, &user.Password)
        if err != nil {
            return nil, err
        }
        // Append each user to the users slice
        users = append(users, user)
    }

    // Check for any error that might have occurred during iteration
    if err = res.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

