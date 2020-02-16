package requestmodels

import (
	"cricket-scoreboard-api/src/models"
)

// PlayerCreateModel godoc
// @Summary Define Player create model
type PlayerCreateModel struct {
	Name       string            `json:"name" form:"name" xml:"name" binding:"required"`
	PlayerType models.PlayerType `json:"playerType" form:"playertype" xml:"playerType" binding:"required"`
}
