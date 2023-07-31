package routes

import (
	"skyzar-assets/logging"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRoutes() {
	Router.GET("/assets/status", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello World"}) })

	Router.NoRoute(func(c *gin.Context) { c.JSON(404, gin.H{"message": "Endpoint Not Found"}) })
	Router.NoMethod(func(c *gin.Context) { c.JSON(405, gin.H{"message": "Method Not Allowed"}) })

	Router.GET("/item/:id", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello World"}) })

	logging.Info("Routes initialized")
}