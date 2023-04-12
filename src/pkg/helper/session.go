package helper

import (
	"crypto/sha1"
	"encoding/base64"
)

func GenerateSessionKey(user User) (string, error) {
	key := RandomStringGenerator(36)

	hasher := sha1.New()
	hasher.Write([]byte(key))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil)), nil

}
