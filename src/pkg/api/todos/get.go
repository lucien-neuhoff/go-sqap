package api_todos

import (
	"net/http"

	db_todos "go-sql/pkg/db/todos"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	todos := db_todos.Get()
	c.IndentedJSON(http.StatusOK, todos)
}

func GetById(c *gin.Context) {
	id := c.Param("id")

	todo, message := db_todos.GetById(id)

	if message != "" {
		c.IndentedJSON(http.StatusBadRequest, message)
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}
