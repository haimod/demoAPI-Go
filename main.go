package main

import (
	"demoAPI-Go/config"
	"demoAPI-Go/models"
	"demoAPI-Go/routes"
	"log"
	"os"
)

func main() {
	// Connect to Supabase PostgreSQL
	config.ConnectDatabase()

	// Auto-migrate: create "users" table if not exists
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
	log.Println("✅ Database migrated successfully!")

	// Setup router
	r := routes.SetupRouter()

	// Get port from environment variable for cloud deployment (Render/Koyeb)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
