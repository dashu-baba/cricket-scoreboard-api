//Package controllers is responsible for returning the response to that request.
package controllers

import (
	"cricket-scoreboard-api/src/requestmodels"
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

//CreateGame creates the game object
func (controller GameController) CreateGame(context *gin.Context) {
	var request requestmodels.GameCreateModel
	var errModel = requestmodels.ValidateCreateTeamsRequests(&request, context)

	if errModel.ErrorCode == http.StatusBadRequest {
		context.JSON(http.StatusBadRequest, errModel)
	}

	controller.GameService.CreateGame(request)

	context.JSON(http.StatusCreated, nil)
}
