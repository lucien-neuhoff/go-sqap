package api_todo

import (
	"net/http"

	db_todo "go-sql/pkg/db/todo"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	todos := db_todo.Get()
	c.IndentedJSON(http.StatusOK, todos)
}

// Todo: edit method to account for user_id
func GetById(c *gin.Context) {
	id := c.Param("id")

	todo, message := db_todo.GetById(id)

	if message != "" {
		c.IndentedJSON(http.StatusBadRequest, message)
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}
