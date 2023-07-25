package routes

import (
	"net/http"

	"skyzar-backend/database"
	"skyzar-backend/logging"

	"github.com/gin-gonic/gin"
)

func GetHypixelReadableNames(c *gin.Context) {
	names, err := database.GetHypixelReadableNames()
	if err != nil {
		logging.Error("Failed to get hypixel readable names, error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, names)
}