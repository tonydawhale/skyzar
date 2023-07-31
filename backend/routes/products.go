package routes

import (
	"skyzar-backend/database"
	"skyzar-backend/logging"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	items, err := database.GetBazaarItemNames()
	if err != nil {
		logging.Error("Failed to get products, error: " + err.Error())
		c.JSON(500, gin.H{"message": "Error fetching products"})
	}
	c.JSON(200, items)
}

func GetProduct(c *gin.Context) {
	item, err := database.GetBazaarItem(c.Param("id"))
	if err != nil {
		logging.Error("Failed to get product, error: " + err.Error())
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(200, item)
}

func GetProductHistory(c *gin.Context) {
	item, err := database.GetBazaarItemHistory(c.Param("id"))
	if err != nil {
		logging.Error("Failed to get product history, error: " + err.Error())
		c.JSON(404, gin.H{"message": "Product history not found"})
		return
	}
	c.JSON(200, item)
}