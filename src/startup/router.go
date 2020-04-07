package startup

import (
	"cricket-scoreboard-api/src/controllers"
	_ "cricket-scoreboard-api/src/docs"
	"cricket-scoreboard-api/src/driver"
	"cricket-scoreboard-api/src/repositories"
	"cricket-scoreboard-api/src/services"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

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

	gameController := controllers.NewGameController(
		services.NewGameService(
			repositories.NewSeriesRepository(
				driver.ConnectDb(),
			),
			repositories.NewTeamRepository(
				driver.ConnectDb(),
			),
			repositories.NewMatchRepository(
				driver.ConnectDb(),
			),
			repositories.NewPlayerRepository(
				driver.ConnectDb(),
			),
		),
	)

	teams := router.Group("/teams")
	{
		teams.GET("", teamController.GetTeams)
		teams.POST("", teamController.CreateTeam)
		teams.GET(":id", teamController.GetTeam)
		teams.PUT(":id", teamController.UpdateTeam)
		teams.POST(":id/players", teamController.AddPlayer)
		teams.DELETE(":id/players/:playerid", teamController.RemovePlayer)
		teams.PUT(":id/players/:playerid", teamController.UpdatePlayer)
	}
	series := router.Group("/series")
	{
		series.POST("", gameController.CreateSeries)
		series.GET(":id", gameController.GetSeries)
		series.PATCH(":id", gameController.UpdateSeriesStatus)
		series.POST(":id/teams", gameController.AddTeams)
		series.DELETE(":id/teams", gameController.RemoveTeams)
		series.PUT(":id/teams/:teamid", gameController.UpdateSquad)
		series.POST(":id/matches", gameController.CreateMatches)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
