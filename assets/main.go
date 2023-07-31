package main

import (
	"skyzar-assets/constants"
	"skyzar-assets/server"
)

func main() {
	constants.LoadConsts()
	server.StartServer()
}