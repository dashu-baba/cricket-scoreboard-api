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

// //GetSeriess ..
// // @Summary Get list of teams
// // @Accept  json
// // @Produce  json
// // @Success 200 {array} responsemodels.Series
// // @Router /teams [get]
// func (controller GameController) GetSeriess(c *gin.Context) {
// 	var (
// 		ctx    context.Context
// 		cancel context.CancelFunc
// 	)

// 	ctx, cancel = context.WithCancel(context.Background())
// 	defer cancel()

// 	c.JSON(http.StatusOK, controller.SeriesService.GetAllSeries(ctx))
// }

//GetSeries ..
// @Summary Get singe item of series
// @Tags Game
// @Accept  json
// @Produce  json
// @Param id path string true "Series ID" string
// @Success 200 {object} responsemodels.Series
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
	}

	res := controller.GameService.CreateSeries(ctx, request)

	c.JSON(http.StatusCreated, res)
}

// //UpdateSeries ...
// // @Summary Update a team item
// // @Accept  json
// // @Produce json
// // @Param model body requestmodels.SeriesUpdateModel true "Update Series"
// // @Param string path int true "Series ID" string
// // @Success 204
// // @Failure 400 {object} responsemodels.ErrorModel
// // @Router /teams/:id [put]
// func (controller GameController) UpdateSeries(c *gin.Context) {
// 	var (
// 		ctx    context.Context
// 		cancel context.CancelFunc
// 	)

// 	ctx, cancel = context.WithCancel(context.Background())
// 	defer cancel()

// 	var request, err = requestmodels.ValidateUpdateSeriessRequests(c)
// 	seriesid := c.Param("id")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
// 			ErrorCode: http.StatusBadRequest,
// 			Message:   err.Error(),
// 		})
// 	}

// 	controller.SeriesService.UpdateSeries(ctx, seriesid, request)

// 	res := responsemodels.Series{
// 		Name:    "Relisource",
// 		Players: make([]responsemodels.Player, 11),
// 	}
// 	c.JSON(http.StatusCreated, res)
// }

//AddTeams ...
// @Summary Add list of teams to the series item
// @Tags Game
// @Accept  json
// @Produce json
// @Param model body requestmodels.TeamsAddRemoveModel true "Add Teams"
// @Param string path int true "Series ID" string
// @Success 201 {object} responsemodels.Series
// @Failure 400 {object} responsemodels.ErrorModel
// @Router /series/:id/teams [post]
func (controller GameController) AddTeams(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	seriesid := c.Param("id")
	var request, err = requestmodels.ValidateAddRemoveTeamRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
	}

	res, err := controller.GameService.AddTeam(ctx, seriesid, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
	}

	c.JSON(http.StatusCreated, res)
}

//RemoveTeams ...
// @Summary Remove list of teams to the series item
// @Tags Game
// @Accept  json
// @Produce json
// @Param model body requestmodels.TeamsAddRemoveModel true "Remove Teams"
// @Param string path int true "Series ID" string
// @Success 204
// @Failure 400 {object} responsemodels.ErrorModel
// @Router /series/:id/teams [delete]
func (controller GameController) RemoveTeams(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	seriesid := c.Param("id")
	var request, err = requestmodels.ValidateAddRemoveTeamRequests(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
	}

	controller.GameService.RemoveTeam(ctx, seriesid, request)

	c.JSON(http.StatusNoContent, nil)
}

// //UpdatePlayer ...
// // @Summary Update a player item
// // @Accept  json
// // @Produce json
// // @Param model body requestmodels.PlayerUpdateModel true "Update Series"
// // @Param id path string true "Series ID" string
// // @Param playerid path string true "Player ID" string
// // @Success 204
// // @Failure 400 {object} responsemodels.ErrorModel
// // @Router /teams/:id/players/:playerid [put]
// func (controller GameController) UpdatePlayer(c *gin.Context) {
// 	var (
// 		ctx    context.Context
// 		cancel context.CancelFunc
// 	)

// 	ctx, cancel = context.WithCancel(context.Background())
// 	defer cancel()

// 	var request, err = requestmodels.ValidateUpdatePlayersRequests(c)
// 	playerid := c.Param("playerid")
// 	seriesid := c.Param("id")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
// 			ErrorCode: http.StatusBadRequest,
// 			Message:   err.Error(),
// 		})
// 	}

// 	controller.SeriesService.UpdatePlayer(ctx, playerid, seriesid, request)

// 	res := responsemodels.Series{
// 		Name:    "Relisource",
// 		Players: make([]responsemodels.Player, 11),
// 	}
// 	c.JSON(http.StatusCreated, res)
// }

// //RemovePlayer ...
// // @Summary Remove a player from the team item
// // @Accept  json
// // @Produce json
// // @Param id path string true "Series ID" string
// // @Param playerid path string true "Player ID" string
// // @Success 204
// // @Router /teams/:id/players/:playerid [delete]
// func (controller GameController) RemovePlayer(c *gin.Context) {
// 	var (
// 		ctx    context.Context
// 		cancel context.CancelFunc
// 	)

// 	ctx, cancel = context.WithCancel(context.Background())
// 	defer cancel()

// 	playerid := c.Param("playerid")
// 	seriesid := c.Param("id")

// 	controller.SeriesService.RemovePlayer(ctx, seriesid, playerid)

// 	c.JSON(http.StatusNoContent, nil)
// }
