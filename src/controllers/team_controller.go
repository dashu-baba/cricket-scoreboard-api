//Package controllers is responsible for returning the response to that request.
package controllers

import (
	"cricket-scoreboard-api/src/models"
	"cricket-scoreboard-api/src/requestmodels"
	"cricket-scoreboard-api/src/responsemodels"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TeamController represents the controller instance.
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

//CreateTeams method returns team lists
func (controller TeamController) CreateTeams(context *gin.Context) {
	var request requestmodels.TeamCreateModel
	var errModel = ValidateCreateTeamsRequests(&request, context)

	if errModel.ErrorCode == http.StatusBadRequest {
		context.JSON(http.StatusBadRequest, errModel)
	}

	res := responsemodels.Team{
		Name:    "Relisource",
		Players: make([]responsemodels.Player, 11),
	}
	for i := 0; i < 11; i++ {
		res.Players[i] = responsemodels.Player{
			Name:       "A",
			PlayerType: models.AllRouner,
		}
	}
	context.JSON(http.StatusCreated, res)
}

//ValidateCreateTeamsRequests method validates the requests payload of CreateTeams method
func ValidateCreateTeamsRequests(TeamCreateModel *requestmodels.TeamCreateModel, c *gin.Context) responsemodels.ErrorModel {
	errorModel := responsemodels.ErrorModel{}
	if err := c.ShouldBind(&TeamCreateModel); err != nil {
		errorModel = responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		}
	}

	return errorModel
}
