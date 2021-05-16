package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"data-pad.app/data-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const (
	// Timeout operations after N seconds
	connectTimeout = 5
)

var (
	client   *mongo.Client
	database *mongo.Database
	ctx      context.Context
)

// GetConnection - Retrieves a client to the MongoDB
func Init() {
	config := config.GetConfig()

	connectionString := config.MongoURI

	parsedConnectionString, err := connstring.ParseAndValidate(connectionString)
	if err != nil {
		log.Printf("Failed to parse connection string: %v", err)
	}

	client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	database = client.Database(parsedConnectionString.Database)
}

func Disconnect() {
	client.Disconnect(ctx)
}

func ClearDB() {
	database.Drop(ctx)
}

func GetDB() *mongo.Database {
	return database
}
