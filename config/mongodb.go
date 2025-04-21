package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() error {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	// Create MongoDB connection string
	uri := fmt.Sprintf("mongodb://%s:%s/%s", host, port, database)
	log.Println(uri)
	//	uri := "mongodb://localhost:27017/lead_crm"
	credential := options.Credential{
		Username: username,
		Password: password,
	}
	// Set client options
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credential)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	Client = client
	fmt.Println("Connected to MongoDB!")

	return nil
}

func GetDB() *mongo.Database {
	return Client.Database(os.Getenv("DB_DATABASE"))
}
