package models

import "database/sql"

type User struct {
	UUID      string       `json:"uuid"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	PublicKey string       `json:"public_key"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type CreateUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	PublicKey string `json:"public_key"`
}
