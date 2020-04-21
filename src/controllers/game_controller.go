//Package controllers is responsible for returning the response to that request.
package controllers

import (
	"context"
	"cricket-scoreboard-api/src/requestmodels/validators"
	"cricket-scoreboard-api/src/responsemodels"
	"cricket-scoreboard-api/src/services"
	"fmt"
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

	var request, err = validators.ValidateCreateSeriesRequests(c)

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
	var request, err = validators.ValidateAddTeamRequests(c)

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
	var request, err = validators.ValidateRemoveTeamRequests(c)

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
	var request, err = validators.ValidateCreateMatchesRequests(c)

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

//GetMatchSummary godoc
// @Summary Get the summary of a match
// @Tags Game
// @Accept  json
// @Produce json
// @Param id path string true "Series ID" string
// @Param matchid path string true "Match ID" string
// @Success 200 {object} responsemodels.MatchSummary
// @Failure 400 {object} responsemodels.ErrorModel
// @Failure 404 {object} responsemodels.ErrorModel
// @Router /series/:id/matches/:matchid [get]
func (controller GameController) GetMatchSummary(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	seriesid := c.Param("id")
	matchid := c.Param("matchid")

	response, errModel := controller.GameService.GetMatchSummary(ctx, seriesid, matchid)

	if errModel != (responsemodels.ErrorModel{}) {
		c.JSON(errModel.ErrorCode, errModel)
		return
	}

	c.JSON(http.StatusOK, response)
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
	var request, err = validators.ValidateUpdateSquadRequests(c)

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
	var request, err = validators.ValidateUpdateSeriesStatusRequests(c)

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

//UpdateMatchStatus godoc
// @Summary Update match status
// @Tags Game
// @Accept  json
// @Produce json
// @Param model body requestmodels.UpdateSeriesStatusModel true "Update Match Status Model"
// @Param id path string true "Series ID" string
// @Param matchid path string true "Match ID" string
// @Success 204
// @Failure 400 {object} responsemodels.ErrorModel
// @Failure 404 {object} responsemodels.ErrorModel
// @Router /series/:id/matches/:matchid [patch]
func (controller GameController) UpdateMatchStatus(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	matchid := c.Param("matchid")
	var request, err = validators.ValidateUpdateSeriesStatusRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	errModel := controller.GameService.UpdateMatchStatus(ctx, matchid, request)
	if errModel != (responsemodels.ErrorModel{}) {
		c.JSON(errModel.ErrorCode, errModel)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

//UpdateMatchPlayingSquad godoc
// @Summary Update starting players of the match
// @Tags Game
// @Accept  json
// @Produce json
// @Param model body requestmodels.MatchPlayingSquadModel true "Match Playing Squad"
// @Param id path string true "Series ID" string
// @Param matchid path string true "Match ID" string
// @Success 204
// @Failure 400 {object} responsemodels.ErrorModel
// @Failure 404 {object} responsemodels.ErrorModel
// @Router /series/:id/matches/:matchid [put]
func (controller GameController) UpdateMatchPlayingSquad(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	seriesid := c.Param("id")
	matchid := c.Param("matchid")
	var request, err = validators.ValidateUpdateMatchPlayingSquadRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	errModel := controller.GameService.UpdateMatchPlayingSquad(ctx, seriesid, matchid, request)

	if errModel != (responsemodels.ErrorModel{}) {
		c.JSON(errModel.ErrorCode, errModel)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

//CreateInnings godoc
// @Summary Created an innings
// @Tags Game
// @Accept  json
// @Produce json
// @Param model body requestmodels.CreateInningsModel true "Create Innings Model"
// @Param id path string true "Series ID" string
// @Param matchid path string true "Match ID" string
// @Success 201 {object} gin.H
// @Failure 400 {object} responsemodels.ErrorModel
// @Failure 404 {object} responsemodels.ErrorModel
// @Router /series/:id/matches/:matchid/innings [post]
func (controller GameController) CreateInnings(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	seriesid := c.Param("id")
	matchid := c.Param("matchid")
	var request, err = validators.ValidateCreateInningsModel(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	id, errModel := controller.GameService.CreateInnings(ctx, seriesid, matchid, request)

	if errModel != (responsemodels.ErrorModel{}) {
		c.JSON(errModel.ErrorCode, errModel)
		return
	}

	url, err := fmt.Printf("/series/%s/matches/%s/innings/%s", seriesid, matchid, id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"resource": url,
	})
}

// //GetInnings godoc
// // @Summary Created an innings
// // @Tags Game
// // @Accept  json
// // @Produce json
// // @Param id path string true "Series ID" string
// // @Param matchid path string true "Match ID" string
// // @Param inningsid path string true "Innings ID" string
// // @Success 200 {object} gin.H
// // @Failure 400 {object} responsemodels.ErrorModel
// // @Failure 404 {object} responsemodels.ErrorModel
// // @Router /series/:id/matches/:matchid/innings/:inningsid [get]
// func (controller GameController) GetInnings(c *gin.Context) {
// 	var (
// 		ctx    context.Context
// 		cancel context.CancelFunc
// 	)

// 	ctx, cancel = context.WithCancel(context.Background())
// 	defer cancel()

// 	seriesid := c.Param("id")
// 	matchid := c.Param("matchid")
// 	var request, err = validators.ValidateCreateInningsModel(c)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
// 			ErrorCode: http.StatusBadRequest,
// 			Message:   err.Error(),
// 		})
// 		return
// 	}

// 	id, errModel := controller.GameService.CreateInnings(ctx, seriesid, matchid, request)

// 	if errModel != (responsemodels.ErrorModel{}) {
// 		c.JSON(errModel.ErrorCode, errModel)
// 		return
// 	}

// 	url, err := fmt.Printf("/series/%s/matches/%s/innings/%s", seriesid, matchid, id)
// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"resource": url,
// 	})
// }
