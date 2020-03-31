package startup

import (
	"cricket-scoreboard-api/src/controllers"
	"cricket-scoreboard-api/src/driver"
	"cricket-scoreboard-api/src/repositories"
	"cricket-scoreboard-api/src/services"

	"github.com/gin-gonic/gin"
)

//NewRouter creates a gin instance and
// returns it.
func NewRouter() *gin.Engine {
	router := gin.New()
	teamController := controllers.NewTeamController(
		services.NewTeamService(
			repositories.NewTeamRepository(
				driver.ConnectDb(),
			),
			repositories.NewPlayerRepository(
				driver.ConnectDb(),
			),
		),
	)
	router.GET("/teams", teamController.GetTeams)
	router.POST("/teams", teamController.CreateTeam)
	router.POST("/teams/:id/players", teamController.AddPlayer)
	router.DELETE("/teams/:id/players/:playerid", teamController.RemovePlayer)

	return router
}
