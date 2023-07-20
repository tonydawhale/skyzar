package main

import (
	"skyzar-backend/constants"
	"skyzar-backend/database"
	"skyzar-backend/server"
	"skyzar-backend/tasks"
)

func main() {
	constants.LoadConsts()
	database.StartClient()
	tasks.StartTasks()
	server.StartServer()
}
