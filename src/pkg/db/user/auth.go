package db_user

import (
	"fmt"
	"go-sql/pkg/helper"
	"log"
)

func Create(user helper.User) string {
	_, err := helper.DB.Query("INSERT INTO users (name, password_hash, email) VALUES ( ?, ?, ?)", user.Name, user.PasswordHash, user.Email)
	if err != nil {
		log.Println("db.user.Create: ", err)
		return fmt.Sprintf("Error while creating user with email '%s'", user.Email)
	}

	user, err = UpdateSession(user)
	if err != nil {
		log.Println("db.user.Create: UpdateSession: ", err)
	}

	return fmt.Sprintf("User with email '%s' created successfully", user.Email)
}

func GetByEmail(email string) (helper.User, error) {
	user := helper.User{Email: email}

	err := helper.DB.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE email='%s'", email)).Scan(&user.ID, &user.Name, &user.PasswordHash, &user.Email, &user.SessionKey, &user.SessionStartedAt)

	if err != nil {
		log.Println("db.user.GetByEmail: ", err)
		return helper.User{}, fmt.Errorf("user with email '%s' not found", email)
	}

	return user, nil
}
