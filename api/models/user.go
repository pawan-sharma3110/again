package models

import "github.com/google/uuid"

type User struct {
	UserId   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
