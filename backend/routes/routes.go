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

	bazaar := Router.Group("/api/bazaar")
	{
		bazaar.GET("/products/:id", GetProduct)
	}

	bazaarProtected := Router.Group("/api/bazaar")//, cloudflareMiddleware())
	{
		bazaarProtected.GET("/products/:id/history", GetProductHistory)
		bazaarProtected.GET("/products", GetProducts)
		bazaarProtected.GET("/crafts", GetCrafts)
		bazaarProtected.GET("/demand", GetDemand)
		bazaarProtected.GET("/margin", GetMargin)
		bazaarProtected.GET("/margin_percent", GetMarginPercent)
		bazaarProtected.GET("/npc_resell", GetNpcResell)
	}

	recipes := Router.Group("/api/recipes", apiKeyMiddleware())
	{
		recipes.GET("/", GetRecipes)
		recipes.POST("/", CreateRecipe)
		recipes.GET("/:id", GetRecipe)
		recipes.PUT("/:id", UpdateRecipe)
		recipes.DELETE("/:id", DeleteRecipe)
	}

	Router.GET("/api/hypixel_readable_item_names", GetHypixelReadableNames)

	logging.Info("Routes initialized")
}

func cloudflareMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfToken := c.Request.Header.Get("cf-turnstile-response")
		if cfToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing cloudflare turnstile header"})
			c.Abort()
			return
		}

		form := url.Values{}
		form.Add("secret", constants.CloudflareTurnstileSecret)
		form.Add("response", cfToken)
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

func apiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiToken := c.Request.Header.Get("Api-Key")
		if apiToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing Api-Key header"})
			c.Abort()
			return
		}
		if apiToken != constants.ApiAuthToken {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}