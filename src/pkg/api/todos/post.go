package api_todos

import (
	db_todos "go-sql/pkg/db/todos"
	"go-sql/pkg/helper"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	id := c.DefaultPostForm("id", "nil")
	title := c.DefaultPostForm("id", "None")
	description := c.DefaultPostForm("id", "")

	if id == "nil" {
		log.Println("api.Post: mising id")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "BadRequest: Mising id"})
		return
	}
	todo := helper.Todo{ID: id, Title: title, Description: description, Complete: 0}

	db_todos.Post(&todo)
}
