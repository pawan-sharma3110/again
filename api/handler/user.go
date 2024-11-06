package handler

import (
	"again/api/database"
	"again/api/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Bad Request: Unable to parse request body", http.StatusBadRequest)
		return
	}
	userId, err := database.InsertUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]uuid.UUID{"userID": userId})
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	userId := r.PathValue("userid")
	id, err := uuid.Parse(userId)
	if err != nil {
		log.Println("Invalid UUID string:", err)
		return
	}
	msg, err := database.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"userID": *msg})
}
