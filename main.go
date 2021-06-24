package main

import (
	"data-pad.app/data-api/config"
	"data-pad.app/data-api/db"
	"data-pad.app/data-api/routes"
)

// @title Data Pad Data Api
// @version 1.0
// @description Data Pad Data api specifications
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email hyurtseven81@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
// @schemes http https
func main() {
	config.Init()

	db.Init()
	defer db.Disconnect()

	routes.Run()
}
