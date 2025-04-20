package main

import (
	"crm-go/config"
	"crm-go/router"
	"log"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to MongoDB
	if err := config.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Initialize Router
	router.Initialize()
}
