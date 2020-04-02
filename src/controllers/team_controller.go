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

//GetTeams method returns team lists
func (controller TeamController) GetTeams(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	c.JSON(http.StatusOK, controller.TeamService.GetAllTeam(ctx))
}

//GetTeam method returns team by id
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

//CreateTeam parses input from request and save the input, returns created team.
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

	controller.TeamService.CreateTeam(ctx, request)

	res := responsemodels.Team{
		Name:    "Relisource",
		Players: make([]responsemodels.Player, 11),
	}
	c.JSON(http.StatusCreated, res)
}

//UpdateTeam parses input from request and update.
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

//AddPlayer adds player into team.
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

//UpdatePlayer parses input from request and update.
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

//RemovePlayer removes player from team.
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
