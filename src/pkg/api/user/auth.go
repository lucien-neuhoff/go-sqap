package api_user

import (
	"log"
	"net/http"
	"strings"

	db_user "go-sql/pkg/db/user"
	"go-sql/pkg/helper"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	email, found := c.Params.Get("email")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"message": "auth/missing-email"})
		log.Print("api.user.SignIn: Missing email")
		return
	}
	email = strings.Replace(email, "%40", "@", 1)

	password_hash, found := c.Params.Get("password")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"message": "auth/missing-password"})
		log.Print("api.user.SignIn: Missing password")
		return
	}

	user, err := db_user.GetByEmail(email)

	if user.PasswordHash != password_hash {
		c.JSON(http.StatusBadRequest, gin.H{"message": "auth/password-mismatch"})
		log.Printf("api.user.SignIn: Password mismatch\n	user.PasswordHash=%s\n	password_hash=%s", user.PasswordHash, password_hash)
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "auth/email-not-found"})
		log.Printf("api.user.SignIn: No user was found with the email '%s'", email)
		return
	}

	db_user.UpdateSession(user)

	c.JSON(http.StatusOK, user)
}

func SignUp(c *gin.Context) {
	name := c.DefaultPostForm("username", "nil")
	password_hash := c.DefaultPostForm("password", "nil")
	email := c.DefaultPostForm("email", "nil")

	if name == "nil" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "auth/missing-username"})
		log.Print("api.user.SignUp: Missing username")
		return
	}
	if password_hash == "nil" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "auth/missing-password"})
		log.Print("api.user.SignUp: Missing password")
		return
	}
	if email == "nil" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "auth/missing-email"})
		log.Print("api.user.SignUp: Missing email")
		return
	}

	user := helper.User{Name: name, PasswordHash: string(password_hash), Email: email}

	msg := db_user.Create(user)

	c.JSON(0, msg)
}
