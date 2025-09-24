package main

import (
	"log"
	"os"

	"github.com/mkumar2307/go-restapi/config"
	"github.com/mkumar2307/go-restapi/models"
	"github.com/mkumar2307/go-restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load env vars (optional: you can use godotenv if needed)
	if os.Getenv("DB_TYPE") == "" {
		log.Fatal("DB_TYPE environment variable is not set")
	}

	// Connect to DB
	config.ConnectDatabase()

	// Auto-migrate models
	config.DB.AutoMigrate(&models.User{})

	// Setup Gin
	r := gin.Default()

	// Register routes
	routes.RegisterUserRoutes(r)

	// Start server
	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
