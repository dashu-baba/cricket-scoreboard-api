//Package controllers is responsible for returning the response to that request.
package controllers

import (
	"cricket-scoreboard-api/src/models"
	"cricket-scoreboard-api/src/requestmodels"
	"cricket-scoreboard-api/src/responsemodels"
	"cricket-scoreboard-api/src/services"
	"fmt"
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
	// teams := make([]responsemodels.Team, 1)
	// teams[0] = responsemodels.Team{
	// 	Name:    "Relisource",
	// 	Players: make([]responsemodels.Player, 11),
	// }
	// for i := 0; i < 11; i++ {
	// 	teams[0].Players[i] = responsemodels.Player{
	// 		Name:       "A",
	// 		PlayerType: models.AllRouner,
	// 	}
	// }
	context.JSON(http.StatusOK, controller.TeamService.GetAllTeam())
}

//CreateTeam parses input from request and save the input, returns created team.
func (controller TeamController) CreateTeam(context *gin.Context) {
	var request requestmodels.TeamCreateModel
	// var errModel = ValidateCreateTeamsRequests(&request, context)

	// if errModel.ErrorCode == http.StatusBadRequest {
	// 	context.JSON(http.StatusBadRequest, errModel)
	// }
	request = requestmodels.TeamCreateModel{
		Name:    "Relisource",
		Logo:    "",
		Players: []requestmodels.PlayerCreateModel{},
	}

	for i := 0; i < 11; i++ {
		item := requestmodels.PlayerCreateModel{
			Name:       fmt.Sprintf("A %d", i),
			PlayerType: models.AllRouner,
		}
		request.Players = append(request.Players, item)
	}

	controller.TeamService.CreateTeam(request)

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
