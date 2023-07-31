package main

import (
	"skyzar-backend/constants"
	"skyzar-backend/database"
	"skyzar-backend/server"
)

func main() {
	constants.LoadConsts()
	database.StartClient()
	server.StartServer()
}
