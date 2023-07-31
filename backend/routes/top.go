package routes

import (
	"net/http"

	"skyzar-backend/constants"
	"skyzar-backend/database"
	"skyzar-backend/logging"

	"github.com/gin-gonic/gin"
)

func GetCrafts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}

func GetDemand(c *gin.Context) {
	items, numItems, err := database.GetTopCategory("buy_volume", constants.BazaarTopDemandQuota)
	if err != nil {
		logging.Error("Failed to get top demand items, error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"items": items, "count": len(items), "totalProducts": numItems })
}

func GetMargin(c *gin.Context) {
	items, numItems, err := database.GetTopCategory("margin", constants.BazaarTopMarginQuota)
	if err != nil {
		logging.Error("Failed to get top margin items, error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"items": items, "count": len(items), "totalProducts": numItems })
}

func GetMarginPercent(c *gin.Context) {
	items, numItems, err := database.GetTopCategory("margin_percent", constants.BazaarTopMarginPercentQuota)
	if err != nil {
		logging.Error("Failed to get top margin percent items, error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"items": items, "count": len(items), "totalProducts": numItems })
}

func GetNpcResell(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}