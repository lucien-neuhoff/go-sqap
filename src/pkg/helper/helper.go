package helper

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"math/rand"
	"time"
)

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

func GenerateSessionKey(user User) (string, error) {
	key := RandomStringGenerator(36)

	hasher := sha1.New()
	hasher.Write([]byte(key))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil)), nil

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
