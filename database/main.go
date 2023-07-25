package main

import (
	"skyzar-database/constants"
	"skyzar-database/database"
	"skyzar-database/tasks"
)

func main() {
	constants.LoadConsts()
	database.StartClient()
	tasks.StartTasks()

	select {}
}
