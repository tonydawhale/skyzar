package routes

import (
	"encoding/json"
	"net/http"
	"net/url"

	"skyzar-backend/constants"
	"skyzar-backend/logging"
	"skyzar-backend/structs"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRoutes() {
	Router.GET("/api/status", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello World"}) })
	Router.GET("/favicon.ico", func (c *gin.Context) { c.Status(http.StatusAccepted) })

	Router.NoRoute(func(c *gin.Context) { c.JSON(404, gin.H{"message": "Not Found"}) })
	Router.NoMethod(func(c *gin.Context) { c.JSON(405, gin.H{"message": "Method Not Allowed"}) })

	bazaar := Router.Group("/api/bazaar", cloudflareMiddleware())
	{
		bazaar.GET("/products", GetProducts)
		bazaar.GET("/products/:id", GetProduct)
		bazaar.GET("/crafts" , GetCrafts)
		bazaar.GET("/demand", GetDemand)
		bazaar.GET("/margin", GetMargin)
		bazaar.GET("/margin_percent", GetMarginPercent)
		bazaar.GET("/npc_resell", GetNpcResell)
	}

	recipes := Router.Group("/api/recipes")
	{
		recipes.GET("/", GetRecipes)
		recipes.POST("/", CreateRecipe)
		recipes.GET("/:id", GetRecipe)
		recipes.PUT("/:id", UpdateRecipe)
		recipes.DELETE("/:id", DeleteRecipe)
	}

	logging.Info("Routes initialized")
}

func cloudflareMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body structs.CloudflarePost
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Cloudflare Error"})
			c.Abort()
			return
		}

		form := url.Values{}
		form.Add("secret", constants.CloudflareTurnstileSecret)
		form.Add("response", body.Cf)
		form.Add("remoteip", c.ClientIP())

		CloudflareResp, err := http.PostForm("https://challenges.cloudflare.com/turnstile/v0/siteverify", form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Cloudflare Error"})
			c.Abort()
			return
		}
		var CloudflareResponse structs.CloudflareRes
		if err := json.NewDecoder(CloudflareResp.Body).Decode(&CloudflareResponse); err != nil {
			logging.Error("Error decoding Cloudflare response, error: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			c.Abort()
			return
		}
		if !CloudflareResponse.Success {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Clouddflare Error", "errors": CloudflareResponse.Errors})
			c.Abort()
			return
		}

		c.Next()
	}
}