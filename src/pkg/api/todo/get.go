package api_todo

import (
	"fmt"
	"log"
	"net/http"

	db_todo "go-sql/pkg/db/todo"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	todos := db_todo.Get()
	c.IndentedJSON(http.StatusOK, todos)
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	user_id := c.Request.Header.Get("User_id")
	session_key := c.Request.Header.Get("Session_key")

	todo, err := db_todo.GetById(id, user_id, session_key)

	if err != nil {
		log.Println("api.todo.GetById: ", err)
		c.IndentedJSON(404, gin.H{"message": fmt.Sprintf("Error while getting todo with id '%s'", id)})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}
