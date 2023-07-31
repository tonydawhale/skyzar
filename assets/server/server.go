package server

import (
	"strconv"

	"skyzar-assets/constants"
	"skyzar-assets/logging"
	"skyzar-assets/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()
}

func StartServer() {
	Router.Use(cors.Default())
	Router.Use(gin.Recovery())
	
	routes.Router = Router
	routes.InitRoutes()

	logging.Info("Server starting on " + strconv.Itoa(constants.Port))
	Router.Run()
}