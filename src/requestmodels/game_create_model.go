package requestmodels

import (
	"cricket-scoreboard-api/src/models"
)

// GameCreateModel godoc
// @Summary Define Game create model
type GameCreateModel struct {
	Name     string          `json:"name" form:"name" xml:"name" binding:"required"`
	GameType models.GameType `json:"gameType" form:"gametype" xml:"gameType" binding:"required"`
	Teams    []string        `json:"teams" form:"teams" xml:"teams" binding:"required"`
}
