package helper

import (
	"database/sql"
	"math/rand"
	"time"
)

var DB *sql.DB
var ENVS map[string]string

// Miliseconds
var SESSION_TIMEOUT = 1_800_000

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

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandomStringGenerator(n int) string {
	bytes := make([]byte, n)
	random := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for i := range bytes {
		bytes[i] = symbols[random.Intn(len(symbols))]
	}
	return string(bytes)
}
