package routes

import (
	"github.com/gin-gonic/gin"
)

func GetCrafts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}

func GetDemand(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}

func GetMargin(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}

func GetMarginPercent(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}

func GetNpcResell(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}