package helper

import "database/sql"

var DB *sql.DB
var ENVS map[string]string

// Seconds
var SESSION_TIMEOUT = 1_800

type Todo struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    int    `json:"complete"`
}

type User struct {
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	PasswordHash     string         `json:"password"`
	Email            string         `json:"email"`
	SessionKey       sql.NullString `json:"session_key"`
	SessionStartedAt sql.NullTime   `json:"session_started_at"`
}
