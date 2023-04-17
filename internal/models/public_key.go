package models

type PublicKey struct {
	UserID string `json:"user_id"`
	Key    string `json:"key"`
}

type PublicKeyRequest struct {
	Email string `json:"email"`
	Key   string `json:"key"`
}
