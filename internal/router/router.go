package router

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func CreateRouter() {
	Router = gin.Default()
}
