package main

import (
	"crm-go/config"
	"crm-go/router"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize Router
	router.Initialize()
}
