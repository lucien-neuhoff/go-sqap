package api_server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	todo "go-sql/pkg/api/todo"
	user "go-sql/pkg/api/user"
	"go-sql/pkg/helper"
)

var router *gin.Engine

func Start(host string, port int) {
	router = gin.Default()

	router.GET("/todos", todo.Get)
	router.GET("/todos/:id", todo.GetById)
	router.POST("/todos", todo.Post)

	router.GET("/users/auth/signin/:email/:password", user.SignIn)
	router.POST("/users/auth/signup", user.SignUp)

	router.GET("/kill", closeDb)

	router.Run(host + ":" + fmt.Sprint(port))
}

func closeDb(c *gin.Context) {
	helper.DB.Close()
}
