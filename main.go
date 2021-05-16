package main

import (
	"datapad-data-api/config"
	"datapad-data-api/db"
	"datapad-data-api/routes"
)

func main() {
	config.Init()

	db.Init()
	defer db.Disconnect()

	routes.Run()
}
