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

// GameController represents the controller instance.
type GameController struct {
	GameService *services.GameService
}

//NewGameController returns a new instance of GameController.
func NewGameController(gameService *services.GameService) *GameController {
	return &GameController{
		GameService: gameService,
	}
}

//GetSeries ..
// @Summary Get singe item of series
// @Tags Game
// @Accept  json
// @Produce  json
// @Param id path string true "Series ID" string
// @Success 200 {array} responsemodels.Team
// @Failure 404 {object} responsemodels.ErrorModel
// @Router /series/:id [get]
func (controller GameController) GetSeries(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	seriesid := c.Param("id")
	c.JSON(http.StatusOK, controller.GameService.GetSeries(ctx, seriesid))
}

//CreateSeries ...
// @Summary Create a series item
// @Tags Game
// @Accept  json
// @Produce  json
// @Param model body requestmodels.SeriesCreateModel true "Create Series"
// @Success 201 {object} responsemodels.Series
// @Failure 400 {object} responsemodels.ErrorModel
// @Router /series [post]
func (controller GameController) CreateSeries(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var request, err = requestmodels.ValidateCreateSeriesRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, errModel := controller.GameService.CreateSeries(ctx, request)

	if errModel != (responsemodels.ErrorModel{}) {
		c.JSON(errModel.ErrorCode, errModel)
		return
	}

	c.JSON(http.StatusCreated, res)
}

//AddTeams ...
// @Summary Add list of teams to the series item
// @Tags Game
// @Accept  json
// @Produce json
// @Param model body requestmodels.TeamsAddModel true "Add Teams"
// @Param id path string true "Series ID" string
// @Success 201 {object} responsemodels.Series
// @Failure 400 {object} responsemodels.ErrorModel
// @Failure 404 {object} responsemodels.ErrorModel
// @Router /series/:id/teams [post]
func (controller GameController) AddTeams(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	seriesid := c.Param("id")
	var request, err = requestmodels.ValidateAddTeamRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, errModel := controller.GameService.AddTeam(ctx, seriesid, request)

	if errModel != (responsemodels.ErrorModel{}) {
		c.JSON(errModel.ErrorCode, errModel)
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

//RemoveTeams ...
// @Summary Remove list of teams to the series item
// @Tags Game
// @Accept  json
// @Produce json
// @Param model body requestmodels.TeamsRemoveModel true "Remove Teams"
// @Param id path string true "Series ID" string
// @Success 204
// @Failure 400 {object} responsemodels.ErrorModel
// @Failure 404 {object} responsemodels.ErrorModel
// @Router /series/:id/teams [delete]
func (controller GameController) RemoveTeams(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	seriesid := c.Param("id")
	var request, err = requestmodels.ValidateRemoveTeamRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	controller.GameService.RemoveTeam(ctx, seriesid, request)

	c.JSON(http.StatusNoContent, nil)
}

//CreateMatches godoc
// @Summary Create list of matches under a series
// @Tags Game
// @Accept  json
// @Produce json
// @Param model body requestmodels.MatchCreateModel true "Create Matches"
// @Param id path string true "Series ID" string
// @Success 204
// @Failure 400 {object} responsemodels.ErrorModel
// @Failure 404 {object} responsemodels.ErrorModel
// @Router /series/:id/matches [post]
func (controller GameController) CreateMatches(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	seriesid := c.Param("id")
	var request, err = requestmodels.ValidateCreateMatchesRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	errModel := controller.GameService.CreateMatches(ctx, seriesid, request)

	if errModel != (responsemodels.ErrorModel{}) {
		c.JSON(errModel.ErrorCode, errModel)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

//UpdateSquad godoc
// @Summary Update
// @Tags Game
// @Accept  json
// @Produce json
// @Param model body requestmodels.UpdateSquadModel true "Update Squad model"
// @Param id path string true "Series ID" string
// @Param teamid path string true "Team ID" string
// @Success 204
// @Failure 400 {object} responsemodels.ErrorModel
// @Failure 404 {object} responsemodels.ErrorModel
// @Router /series/:id/teams/:teamid [put]
func (controller GameController) UpdateSquad(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	seriesid := c.Param("id")
	teamid := c.Param("teamid")
	var request, err = requestmodels.ValidateUpdateSquadRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	errModel := controller.GameService.UpdateSquad(ctx, seriesid, teamid, request)

	if errModel != (responsemodels.ErrorModel{}) {
		c.JSON(errModel.ErrorCode, errModel)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

//UpdateSeriesStatus godoc
// @Summary Update series status
// @Tags Game
// @Accept  json
// @Produce json
// @Param model body requestmodels.UpdateSeriesStatusModel true "Update Series Status Model"
// @Param id path string true "Series ID" string
// @Success 204
// @Failure 400 {object} responsemodels.ErrorModel
// @Failure 404 {object} responsemodels.ErrorModel
// @Router /series/:id [patch]
func (controller GameController) UpdateSeriesStatus(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	seriesid := c.Param("id")
	var request, err = requestmodels.ValidateUpdateSeriesStatusRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	errModel := controller.GameService.UpdateSeriesStatus(ctx, seriesid, request)
	if errModel != (responsemodels.ErrorModel{}) {
		c.JSON(errModel.ErrorCode, errModel)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
