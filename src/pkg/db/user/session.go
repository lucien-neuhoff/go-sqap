package db_user

import (
	"database/sql"
	"errors"
	"fmt"
	"go-sql/pkg/helper"
	"log"
	"time"
)

func UpdateSession(user helper.User) (updated_user helper.User, err error) {
	// Session still up
	if !SessionTimedOut(user.SessionStartedAt.Time) {
		return user, nil
	}

	log.Printf("db.user.UpdateSession: Session timed out for user '%s'", user.Name)

	// Session key valid
	valid, err := VerifySessionKey(user.ID, user.SessionKey.String)
	if err != nil {
		return user, err
	}

	if valid {
		return user, nil
	}

	log.Printf("db.user.UpdateSession: Session key mismatch, generating a new one for user '%s'", user.Name)

	session_key, err := helper.GenerateSessionKey(user)
	if err != nil {
		log.Println("db.user.UpdateSession: ", err)
	}

	now := time.Now()

	query, err := helper.DB.Prepare("INSERT INTO users (session_key, session_started_at) VALUES (?, ?)")
	if err != nil {
		log.Println("db.user.UpdateSession: ", err)
	}
	query.Exec(&session_key, now)

	user.SessionKey.String = session_key
	user.SessionStartedAt.Time = now

	return user, nil
}

func GetSessionKey(user helper.User) (found bool, msg sql.NullString) {
	var session_key sql.NullString

	row := helper.DB.QueryRow(fmt.Sprintf("SELECT session_key FROM users WHERE id='%s' AND email='%s", user.ID, user.Email))
	err := row.Scan(&user.SessionKey)
	if err != nil {
		log.Println("db.user.GetSessionKey: ", err)
	}

	return true, session_key
}

func SessionTimedOut(session_started_at time.Time) bool {
	return (time.Now().UTC().UnixMilli() - session_started_at.UnixMilli()) > int64(helper.SESSION_TIMEOUT)
}

func VerifySessionKey(user_id string, session_key string) (valid bool, err error) {
	var user_session_key string

	helper.DB.QueryRow("SELECT session_key FROM users WHERE id=?", user_id).Scan(&user_session_key)

	if user_session_key == session_key {
		return true, nil
	}

	log.Printf("db.user.VerifySessionKey: \n	user_session_key=%s\n	session_key=%s", user_session_key, session_key)

	return false, errors.New("session_key mismatch")
}
