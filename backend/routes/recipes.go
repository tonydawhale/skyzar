package routes

import (
	"net/http"

	"skyzar-backend/database"
	"skyzar-backend/logging"
	"skyzar-backend/structs"

	"github.com/gin-gonic/gin"
)

func GetRecipe(c *gin.Context) {
	recipe, err := database.GetRecipe(c.Param("id"))
	if err != nil {
		logging.Error("Failed to get recipe, error: " + err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

func GetRecipes(c *gin.Context) {
	recipes, err := database.GetRecipes()
	if err != nil {
		logging.Error("Failed to get recipes, error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

func CreateRecipe(c *gin.Context) {
	var recipe structs.SkyblockItemRecipe
	if err := c.BindJSON(&recipe); err != nil {
		logging.Error("Failed to bind JSON, error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}
	if err := database.CreateRecipe(recipe); err != nil {
		logging.Error("Failed to create recipe, error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.Status(http.StatusCreated)
}

func UpdateRecipe(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}

func DeleteRecipe(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}