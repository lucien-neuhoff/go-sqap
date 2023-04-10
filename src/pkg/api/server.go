package api_server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	todos "go-sql/pkg/api/todos"
	"go-sql/pkg/helper"
)

var router *gin.Engine

func Start(host string, port int) {
	router = gin.Default()

	router.GET("/todos", todos.Get)
	router.GET("/todos/:id", todos.GetById)

	router.POST("/todos", todos.Post)

	router.GET("/kill", closeDb)

	router.Run(host + ":" + fmt.Sprint(port))
}

func closeDb(c *gin.Context) {
	helper.DB.Close()
}
