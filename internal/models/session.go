package models

import (
	"crypto/rsa"
	"time"
)

type Session struct {
	PublicKey rsa.PublicKey `json:"public_key"`
	UserID    *string       `json:"user_id"`
	Token     string        `json:"token"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type CreateSessionRequest struct {
	PublicKey string  `json:"public_key"`
	UserID    *string `json:"user_id"`
}
