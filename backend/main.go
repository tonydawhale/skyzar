package main

import (
	"skyzar-backend/constants"
	"skyzar-backend/server"
	"skyzar-backend/tasks"
)

func main() {
	constants.LoadConsts()
	tasks.StartTasks()
	server.StartServer()
}
