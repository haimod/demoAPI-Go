package routes

import (
	"demoAPI-Go/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures all API routes and returns the Gin engine
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// API v1 group
	api := r.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.GET("/users/:id", handlers.GetUserByID)
		api.POST("/users", handlers.CreateUser)
		api.DELETE("/users/:id", handlers.DeleteUser)
	}

	return r
}
