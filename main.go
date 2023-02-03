package main

import (
	"contentManagerBackend/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	fmt.Println("Hello, World!")
	// Define routes and their respective handlers
	router.POST("/signup", handler.Signup)

	// Start the Gin server
	router.Run()
}

