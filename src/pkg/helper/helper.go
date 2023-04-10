package helper

import "database/sql"

var DB *sql.DB
var ENVS map[string]string

type Todo struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    int    `json:"complete"`
}

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"password"`
	Email        string `json:"email"`
}

type Message struct {
	From    string `json:"from"`
	Content string `json:"message"`
}
