// Package config is application configuration module
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	AppName  string
	MongoURI string
}

// Config - exported struct to be used anywhere
var config appConfig

// Init initializes config
func InitWithEnvFile(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config = appConfig{
		AppName:  os.Getenv("APP_NAME"),
		MongoURI: os.Getenv("MONGO_URI"),
	}

}

// Init initializes config
func Init() {
	InitWithEnvFile(".env")
}

func GetConfig() appConfig {
	return config
}
