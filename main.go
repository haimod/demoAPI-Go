package main

import (
	"demoAPI-Go/routes"
	"log"
	"os"
)

func main() {
	r := routes.SetupRouter()

	// Get port from environment variable for cloud deployment (Render/Koyeb)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port for local development
	}

	log.Printf("🚀 Server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
