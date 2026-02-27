package handlers

import (
	"demoAPI-Go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers returns all users
// GET /api/users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    models.Users,
	})
}

// GetUserByID returns a single user by ID
// GET /api/users/:id
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	for _, user := range models.Users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"data":    user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// CreateUser adds a new user
// POST /api/users
func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Auto-increment ID
	if len(models.Users) > 0 {
		newUser.ID = models.Users[len(models.Users)-1].ID + 1
	} else {
		newUser.ID = 1
	}

	models.Users = append(models.Users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    newUser,
	})
}

// DeleteUser removes a user by ID
// DELETE /api/users/:id
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	for i, user := range models.Users {
		if user.ID == id {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
