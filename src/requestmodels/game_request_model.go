package requestmodels

import (
	"cricket-scoreboard-api/src/models"
	"errors"

	"github.com/gin-gonic/gin"
)

// SeriesCreateModel godoc
// @Summary Define Game create model
type SeriesCreateModel struct {
	Name     string          `json:"name" form:"name" xml:"name" binding:"required"`
	GameType models.GameType `json:"gameType" form:"gametype" xml:"gameType" binding:"required"`
	Teams    []string        `json:"teams" form:"teams" xml:"teams" binding:"required"`
}

// TeamsAddRemoveModel godoc
// @Summary Define team add or remove model
type TeamsAddRemoveModel struct {
	Teams []string `json:"teams" form:"teams" xml:"teams" binding:"required"`
}

//ValidateCreateSeriesRequests method validates the requests payload of CreateSeries method
func ValidateCreateSeriesRequests(c *gin.Context) (SeriesCreateModel, error) {
	model := SeriesCreateModel{}
	model.Teams = []string{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	if model.GameType == models.Bilateral && len(model.Teams) > 2 {
		return model, errors.New("Invalid data, bilateral series only have max 2 teams")
	}

	return model, nil
}

//ValidateAddRemoveTeamRequests method validates the requests payload of AddTeams and RemoveTeams method
func ValidateAddRemoveTeamRequests(c *gin.Context) (TeamsAddRemoveModel, error) {
	model := TeamsAddRemoveModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}
