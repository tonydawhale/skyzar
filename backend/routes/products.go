package routes

import (
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}

func GetProduct(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}