package db_user

import (
	"database/sql"
	"fmt"
	"go-sql/pkg/helper"
	"log"
	"time"
)

func UpdateSession(user helper.User) error {
	if !SessionTimedOut(user.SessionStartedAt.Time) {
		return nil
	}

	session_key, err := helper.GenerateSessionKey(user)
	if err != nil {
		log.Panic("db.user.UpdateSession: ", err)
	}

	query, err := helper.DB.Prepare("INSERT INTO users (session_key, session_started_at) VALUES (?, ?)")
	if err != nil {
		log.Panic("db.user.UpdateSession: ", err)
	}
	query.Exec(&session_key, time.Now())

	// // Updates the session_key
	// found, session_key := GetSessionKey(user)
	// if !found {
	// 	msg := fmt.Sprintf("SessionKey not found for user with id '%s'", user.ID)
	// 	log.Println("db.user.UpdateSession", msg)
	// 	return msg
	// }

	return nil
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
