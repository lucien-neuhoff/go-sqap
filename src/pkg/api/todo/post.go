package api_todo

import (
	"go-sql/pkg/helper"
	"log"
	"net/http"

	db_todo "go-sql/pkg/db/todo"
	db_user "go-sql/pkg/db/user"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	id := c.DefaultPostForm("id", "0")
	title := c.DefaultPostForm("title", "None")
	description := c.DefaultPostForm("description", "")

	user_id := c.Request.Header.Get("User_id")
	session_key := c.Request.Header.Get("Session_key")

	user := helper.User{ID: user_id}
	user.SessionKey.String = session_key

	valid, err := db_user.VerifySessionKey(user.ID, user.SessionKey.String)
	if err != nil {
		log.Printf("api.todo.Post: %s", err)
	}

	if valid {
		log.Printf("api.todo.Post: valid session_key for user_id=%s", user_id)
	}

	if user_id == "" {
		log.Println("api.todo.Post: Missing user_id")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "BadRequest: Missing user_id"})
		return
	}

	todo := helper.Todo{ID: id, UserID: user_id, Title: title, Description: description, Complete: 0}

	db_todo.Post(&todo)
}
