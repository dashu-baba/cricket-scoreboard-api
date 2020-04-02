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

// TeamUpdateModel godoc
// @Summary Define Team update model
type TeamUpdateModel struct {
	Name string `json:"name" form:"name" xml:"name" binding:"required"`
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
}

// PlayerUpdateModel godoc
// @Summary Define Player update model
type PlayerUpdateModel struct {
	Name       string            `json:"name" form:"name" xml:"name" binding:"required"`
	PlayerType models.PlayerType `json:"playerType" form:"playertype" xml:"playerType" binding:"required"`
}

//ValidateUpdateTeamsRequests method validates the requests payload of UpdateTeam method
func ValidateUpdateTeamsRequests(c *gin.Context) (TeamUpdateModel, error) {
	var model TeamUpdateModel
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}

//ValidateUpdatePlayersRequests method validates the requests payload of UpdatePlayer method
func ValidateUpdatePlayersRequests(c *gin.Context) (PlayerUpdateModel, error) {
	var model PlayerUpdateModel
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}

//ValidateCreateTeamsRequests method validates the requests payload of CreateTeams method
func ValidateCreateTeamsRequests(c *gin.Context) (TeamCreateModel, error) {
	model := TeamCreateModel{}
	model.Players = []PlayerModel{}
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
