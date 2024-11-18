package main

import (
	"hospital-middleware/config"
	"hospital-middleware/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, db)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}