package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ServerConnection holds the MongoDB client
type ServerConnection struct {
	Client *mongo.Client
}

// Connect_could establishes a connection to the MongoDB server
func (SC *ServerConnection) Connect_could() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env", err.Error())
	}

	// Get the database URL from environment variables
	url := os.Getenv("DB_URL")

	// Set server API options for MongoDB client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	options := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	// Connect to the MongoDB server
	client, connetion_err := mongo.Connect(context.TODO(), options)
	if connetion_err != nil {
		log.Panic("Failed to connect to server", connetion_err.Error())
	}

	// Ping the database to verify the connection
	if err := client.Database("BlogPost").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Panic("Ping failed", err.Error())
	}

	// Assign the connected client to the ServerConnection struct
	SC.Client = client
	log.Println("Connected to server")
}
