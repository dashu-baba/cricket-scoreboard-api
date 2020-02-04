package controllers

import (
	"cricket-scoreboard-api/src/models"
	"cricket-scoreboard-api/src/responsemodels"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TeamController represents the controller instance
type TeamController struct{}

//GetTeams method returns team lists
func (controller TeamController) GetTeams(context *gin.Context) {
	teams := make([]responsemodels.Team, 1)
	teams[0] = responsemodels.Team{
		Name:    "Relisource",
		Players: make([]responsemodels.Player, 11),
	}
	for i := 0; i < 11; i++ {
		teams[0].Players[i] = responsemodels.Player{
			Name:       "A",
			PlayerType: models.AllRouner,
		}
	}
	context.JSON(http.StatusOK, teams)
}
