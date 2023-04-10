package api_todo

import (
	"go-sql/pkg/helper"
	"log"
	"net/http"

	db_todo "go-sql/pkg/db/todo"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	user_id := c.DefaultPostForm("user_id", "nil")
	title := c.DefaultPostForm("title", "None")
	description := c.DefaultPostForm("description", "")

	if user_id == "nil" {
		log.Println("api.todo.Post: Missing user_id")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "BadRequest: Missing user_id"})
		return
	}
	todo := helper.Todo{UserID: user_id, Title: title, Description: description, Complete: 0}

	db_todo.Post(&todo)
}
