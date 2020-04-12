package main

import (
	"cricket-scoreboard-api/src/models"
	"cricket-scoreboard-api/src/startup"
)

// @title Cricket Scoreboard API
// @version 1.0
// @description This contains the commn REST api to support a modern cricket scoreboard project.
// @termsOfService https://github.com/dashu-baba/cricket-scoreboard-api
// @license.name MIT
// @license.url https://github.com/dashu-baba/cricket-scoreboard-api/blob/master/LICENSE
func main() {
	configuration := models.Configuration

	// Configuring port
	port := configuration.Server.Port
	if port == "" {
		port = "8080"
	}

	// Setting New Router
	router := startup.NewRouter()
	_ = router.Run(":" + port)
}
