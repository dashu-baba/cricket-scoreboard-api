//Package controllers is responsible for returning the response to that request.
package controllers

import (
	"context"
	"cricket-scoreboard-api/src/requestmodels/validators"
	"cricket-scoreboard-api/src/responsemodels"
	"cricket-scoreboard-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InningsController represents the controller instance.
type InningsController struct {
	InningsService *services.InningsService
}

//NewInningsController returns a new instance of InningsController.
func NewInningsController(inningsService *services.InningsService) *InningsController {
	return &InningsController{
		InningsService: inningsService,
	}
}

//UpdateOver ..
// @Summary Update an over
// @Tags Innings
// @Accept  json
// @Produce  json
// @Param model body requestmodels.OverUpdateModel true "Over Update Model"
// @Param inningsid path string true "Innings ID" string
// @Param overid path string true "Over ID" string
// @Success 204
// @Failure 404 {object} responsemodels.ErrorModel
// @Failure 400 {object} responsemodels.ErrorModel
// @Router /innings/:inningsid/over/:overid [put]
func (controller InningsController) UpdateOver(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	overid := c.Param("overid")
	inningsid := c.Param("inningsid")
	var request, err = validators.ValidateOverUpdateModel(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	controller.InningsService.UpdateOver(ctx, inningsid, overid, request)

	c.JSON(http.StatusNoContent, nil)
}

//StartInnings godoc
// @Summary Created an innings
// @Tags Innings
// @Accept  json
// @Produce json
// @Param model body requestmodels.StartInningsModel true "Start Innings Model"
// @Param id path string true "Series ID" string
// @Param matchid path string true "Match ID" string
// @Param inningsid path string true "Innings ID" string
// @Success 204
// @Failure 400 {object} responsemodels.ErrorModel
// @Failure 404 {object} responsemodels.ErrorModel
// @Router /innings/:inningsid/start [put]
func (controller InningsController) StartInnings(c *gin.Context) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	seriesid := c.Param("id")
	matchid := c.Param("matchid")
	inningsid := c.Param("inningsid")
	var request, err = validators.ValidateStartInningsModel(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	controller.InningsService.StartInnings(ctx, seriesid, matchid, inningsid, request)

	c.JSON(http.StatusNoContent, nil)
}
