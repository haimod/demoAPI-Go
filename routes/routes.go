package routes

import (
	"demoAPI-Go/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures all API routes and returns the Gin engine
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS middleware - cho phép mọi origin gọi API
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	// API v1 group
	api := r.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.GET("/users/:id", handlers.GetUserByID)
		api.POST("/users", handlers.CreateUser)
		api.PUT("/users/:id", handlers.UpdateUser)
		api.DELETE("/users/:id", handlers.DeleteUser)
	}

	return r
}
