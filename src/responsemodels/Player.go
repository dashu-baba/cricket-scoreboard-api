package responsemodels

import (
	"cricket-scoreboard-api/src/models"
)

// Player godoc
// @Summary Define Player model
type Player struct {
	ID         string `json:"id" form:"id" xml:"id"`
	Name       string `json:"name" form:"name" xml:"name"`
	PlayerType models.PlayerType `json:"playerType" form:"playertype" xml:"playerType"`
	TeamID     string `json:"teamID" form:"teamid" xml:"teamID"`
}
