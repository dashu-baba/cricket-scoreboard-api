package main

import (
	"cricket-scoreboard-api/src/models"
	"cricket-scoreboard-api/src/startup"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variable from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading env file %v", err)
	}

	configuration := models.Configuration
	router := startup.NewRouter()
	router.Run(":" + configuration.Server.Port)
}
