package test

import (
	"data-pad.app/data-api/config"
	"data-pad.app/data-api/db"

	"github.com/gin-gonic/gin"
)

func Init() {
	gin.SetMode(gin.TestMode)

	config.InitWithEnvFile("../../test.env")
	db.Init()

	defer db.ClearDB()

	defer db.Disconnect()
}
