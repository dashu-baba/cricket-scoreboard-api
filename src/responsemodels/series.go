package responsemodels

import (
	"cricket-scoreboard-api/src/models"
)

// Series godoc
// @Summary Define Series model
type Series struct {
	ID     string             `json:"id" form:"id" xml:"id"`
	Name   string             `json:"name" form:"name" xml:"name"`
	Type   models.GameType    `json:"type" form:"type" xml:"type"`
	Teams  []Team             `json:"teams" form:"teams" xml:"teams"`
	Status models.SeriesState `json:"status" form:"status" xml:"status"`
}
