package test

import (
	"datapad-data-api/config"
	"datapad-data-api/db"

	"github.com/gin-gonic/gin"
)

func Init() {
	gin.SetMode(gin.TestMode)

	config.InitWithEnvFile("../../test.env")
	db.Init()

	defer db.ClearDB()

	defer db.Disconnect()
}
