package handler

import (
	"contentManagerBackend/model"
	"contentManagerBackend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	// Get user information from the request
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the Signup method in the service package
	if err := service.Signup(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return success
	c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
}
