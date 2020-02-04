package main

import (
	"cricket-scoreboard-api/src/models"
	"cricket-scoreboard-api/src/startup"
)

func main() {
	configuration := models.Configuration
	router := startup.NewRouter()
	router.Run(":" + configuration.Server.Port)
}
