package startup

import (
	_ "cricket-scoreboard-api/src/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"cricket-scoreboard-api/src/controllers"
	"cricket-scoreboard-api/src/driver"
	"cricket-scoreboard-api/src/repositories"
	"cricket-scoreboard-api/src/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//NewRouter creates a gin instance and
// returns it.
func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/teams", teamController.GetTeams)
	router.POST("/teams", teamController.CreateTeam)
	router.GET("/teams/:id", teamController.GetTeam)
	router.PUT("/teams/:id", teamController.UpdateTeam)
	router.POST("/teams/:id/players", teamController.AddPlayer)
	router.DELETE("/teams/:id/players/:playerid", teamController.RemovePlayer)
	router.PUT("/teams/:id/players/:playerid", teamController.UpdatePlayer)

	return router
}
