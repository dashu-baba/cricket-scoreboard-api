//Package controllers is responsible for returning the response to that request.
package controllers

import (
	"cricket-scoreboard-api/src/requestmodels"
	"cricket-scoreboard-api/src/responsemodels"
	"cricket-scoreboard-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TeamController represents the controller instance.
type TeamController struct {
	TeamService *services.TeamService
}

//NewTeamController returns a new instance of TeamController.
func NewTeamController(TeamService *services.TeamService) *TeamController {
	return &TeamController{
		TeamService: TeamService,
	}
}

//GetTeams method returns team lists
func (controller TeamController) GetTeams(context *gin.Context) {
	context.JSON(http.StatusOK, controller.TeamService.GetAllTeam())
}

//CreateTeam parses input from request and save the input, returns created team.
func (controller TeamController) CreateTeam(context *gin.Context) {
	var request, err = requestmodels.ValidateCreateTeamsRequests(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
	}

	controller.TeamService.CreateTeam(request)

	res := responsemodels.Team{
		Name:    "Relisource",
		Players: make([]responsemodels.Player, 11),
	}
	context.JSON(http.StatusCreated, res)
}

//AddPlayer adds player into team.
func (controller TeamController) AddPlayer(context *gin.Context) {
	var request, err = requestmodels.ValidateAddPlayerRequests(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
	}

	res := controller.TeamService.CreatePlayer(request)

	context.JSON(http.StatusCreated, res)
}

//RemovePlayer removes player from team.
func (controller TeamController) RemovePlayer(context *gin.Context) {
	playerid := context.Param("playerid")
	teamid := context.Param("id")

	controller.TeamService.RemovePlayer(teamid, playerid)

	context.JSON(http.StatusNoContent, nil)
}

//UpdateTeam update team partially, and return no content.
func (controller TeamController) UpdateTeam(context *gin.Context) {

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
