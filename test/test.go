package test

import (
	"data-pad.app/data-api/config"
	"data-pad.app/data-api/db"

	"github.com/gin-gonic/gin"
)

func Init(env_file string) {
	gin.SetMode(gin.TestMode)

	if env_file == "" {
		env_file = "../../test.env"
	}

	config.InitWithEnvFile(env_file)
	db.Init()
}

func Clear() {
	db.ClearDB()

	db.Disconnect()
}
