package requestmodels

import (
	"cricket-scoreboard-api/src/models"

	"github.com/gin-gonic/gin"
)

// TeamCreateModel godoc
// @Summary Define Team create model
type TeamCreateModel struct {
	Name    string        `json:"name" form:"name" xml:"name" binding:"required"`
	Logo    string        `json:"logo" form:"logo" xml:"logo"`
	Players []PlayerModel `json:"players" form:"players" xml:"players"`
}

// PlayerModel godoc
// @Summary Define Players within Team create model
type PlayerModel struct {
	Name       string            `json:"name" form:"name" xml:"name" binding:"required"`
	PlayerType models.PlayerType `json:"playerType" form:"playertype" xml:"playerType" binding:"required"`
}

// PlayerCreateModel godoc
// @Summary Define Player create model
type PlayerCreateModel struct {
	Name       string            `json:"name" form:"name" xml:"name" binding:"required"`
	PlayerType models.PlayerType `json:"playerType" form:"playertype" xml:"playerType" binding:"required"`
	TeamID     string            `json:"teamId" form:"teamid" xml:"teamId" binding:"required"`
}

//ValidateCreateTeamsRequests method validates the requests payload of CreateTeams method
func ValidateCreateTeamsRequests(c *gin.Context) (TeamCreateModel, error) {
	var model TeamCreateModel
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}

//ValidateAddPlayerRequests method validates the requests payload of AddPlayer method
func ValidateAddPlayerRequests(c *gin.Context) (PlayerCreateModel, error) {
	var model PlayerCreateModel
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}
