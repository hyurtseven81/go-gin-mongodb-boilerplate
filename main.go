package main

import (
	"data-pad.app/data-api/config"
	"data-pad.app/data-api/db"
	"data-pad.app/data-api/routes"
)

func main() {
	config.Init()

	db.Init()
	defer db.Disconnect()

	routes.Run()
}
