//Package controllers is responsible for returning the response to that request.
package controllers

import (
	"context"
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

//GetTeams ..
// @Summary Get list of teams
// @Accept  json
// @Produce  json
// @Success 200 {array} responsemodels.Team
// @Router /teams [get]
func (controller TeamController) GetTeams(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	c.JSON(http.StatusOK, controller.TeamService.GetAllTeam(ctx))
}

//GetTeam ..
// @Summary Get singe item of team
// @Accept  json
// @Produce  json
// @Param id path string true "Team ID" string
// @Success 200 {object} responsemodels.Team
// @Router /teams/:id [get]
func (controller TeamController) GetTeam(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	teamid := c.Param("id")
	c.JSON(http.StatusOK, controller.TeamService.GetTeam(ctx, teamid))
}

//CreateTeam ...
// @Summary Create a team item
// @Accept  json
// @Produce  json
// @Param model body requestmodels.TeamCreateModel true "Create Team"
// @Success 201 {object} responsemodels.Team
// @Failure 400 {object} responsemodels.ErrorModel
// @Router /teams [post]
func (controller TeamController) CreateTeam(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var request, err = requestmodels.ValidateCreateTeamsRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
	}

	res := controller.TeamService.CreateTeam(ctx, request)

	c.JSON(http.StatusCreated, res)
}

//UpdateTeam ...
// @Summary Update a team item
// @Accept  json
// @Produce json
// @Param model body requestmodels.TeamUpdateModel true "Update Team"
// @Param string path int true "Team ID" string
// @Success 204 
// @Failure 400 {object} responsemodels.ErrorModel
// @Router /teams/:id [put]
func (controller TeamController) UpdateTeam(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var request, err = requestmodels.ValidateUpdateTeamsRequests(c)
	teamid := c.Param("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
	}

	controller.TeamService.UpdateTeam(ctx, teamid, request)

	res := responsemodels.Team{
		Name:    "Relisource",
		Players: make([]responsemodels.Player, 11),
	}
	c.JSON(http.StatusCreated, res)
}

//AddPlayer ...
// @Summary Add a player to the team item
// @Accept  json
// @Produce json
// @Param model body requestmodels.PlayerCreateModel true "Add Team"
// @Param string path int true "Team ID" string
// @Success 201 {object} responsemodels.Player 
// @Failure 400 {object} responsemodels.ErrorModel
// @Router /teams/:id/players [post]
func (controller TeamController) AddPlayer(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	teamid := c.Param("id")
	var request, err = requestmodels.ValidateAddPlayerRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
	}

	res := controller.TeamService.CreatePlayer(ctx, teamid, request)

	c.JSON(http.StatusCreated, res)
}

//UpdatePlayer ...
// @Summary Update a player item
// @Accept  json
// @Produce json
// @Param model body requestmodels.PlayerUpdateModel true "Update Team"
// @Param id path string true "Team ID" string
// @Param playerid path string true "Player ID" string
// @Success 204
// @Failure 400 {object} responsemodels.ErrorModel
// @Router /teams/:id/players/:playerid [put]
func (controller TeamController) UpdatePlayer(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var request, err = requestmodels.ValidateUpdatePlayersRequests(c)
	playerid := c.Param("playerid")
	teamid := c.Param("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
	}

	controller.TeamService.UpdatePlayer(ctx, playerid, teamid, request)

	res := responsemodels.Team{
		Name:    "Relisource",
		Players: make([]responsemodels.Player, 11),
	}
	c.JSON(http.StatusCreated, res)
}

//RemovePlayer ...
// @Summary Remove a player from the team item
// @Accept  json
// @Produce json
// @Param id path string true "Team ID" string
// @Param playerid path string true "Player ID" string
// @Success 204
// @Router /teams/:id/players/:playerid [delete]
func (controller TeamController) RemovePlayer(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	playerid := c.Param("playerid")
	teamid := c.Param("id")

	controller.TeamService.RemovePlayer(ctx, teamid, playerid)

	c.JSON(http.StatusNoContent, nil)
}
