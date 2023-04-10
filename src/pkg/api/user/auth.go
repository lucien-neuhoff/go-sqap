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
		msg := "api.user.SignIn: Missing email"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	email = strings.Replace(email, "%40", "@", 1)

	password_hash, found := c.Params.Get("password")
	if !found {
		msg := "api.user.SignIn: Missing password"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	user, msg := db_user.GetByEmail(email)

	if user.PasswordHash != password_hash {
		log.Println("api.user.SignIn: Passwords do not match")
		c.JSON(http.StatusBadRequest, "Passwords do not match")
		return
	}

	if msg != "" {
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	c.JSON(http.StatusOK, user)
}

func SignUp(c *gin.Context) {
	name := c.DefaultPostForm("name", "nil")
	// Password Hashing has to be done client side
	password_hash := c.DefaultPostForm("password", "nil")
	email := c.DefaultPostForm("email", "nil")

	if name == "nil" {
		msg := "api.user.SignUp: Missing name"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	if password_hash == "nil" {
		msg := "api.user.SignUp: Missing password"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	if email == "nil" {
		msg := "api.user.SignUp: Missing email"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	user := helper.User{Name: name, PasswordHash: string(password_hash), Email: email}

	msg := db_user.SignUp(user)

	c.JSON(0, msg)
}
