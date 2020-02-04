package startup

import (
	"cricket-scoreboard-api/src/controllers"

	"github.com/gin-gonic/gin"
)

//NewRouter creates a gin instance and
// returns it.
func NewRouter() *gin.Engine {
	router := gin.New()

	teams := new(controllers.TeamController)
	router.GET("/teams", teams.GetTeams)

	return router
}
