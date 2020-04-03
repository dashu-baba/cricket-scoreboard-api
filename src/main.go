package main

import (
	"cricket-scoreboard-api/src/models"
	"cricket-scoreboard-api/src/startup"
	"log"

	"github.com/joho/godotenv"
)

// @title Cricket Scoreboard API
// @version 1.0
// @description This contains the commn REST api to support a modern cricket scoreboard project.
// @termsOfService https://github.com/dashu-baba/cricket-scoreboard-api
// @license.name MIT
// @license.url https://github.com/dashu-baba/cricket-scoreboard-api/blob/master/LICENSE
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
