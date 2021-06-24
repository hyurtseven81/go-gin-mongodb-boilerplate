// Package config is application configuration module
package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	AppName  string
	Host     string
	MongoURI string
}

// Config - exported struct to be used anywhere
var config appConfig

// Init initializes config
func InitWithEnvFile(envFile string) {
	flag.Parse()

	err := godotenv.Load(envFile)
	if err != nil {
		log.Print("Error loading .env file")
	}

	config = appConfig{
		AppName:  os.Getenv("APP_NAME"),
		Host:     os.Getenv("HOST"),
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
