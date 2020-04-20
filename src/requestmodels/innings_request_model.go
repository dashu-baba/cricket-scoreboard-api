package requestmodels

import (
	"cricket-scoreboard-api/src/models"
)

// OverUpdateModel godoc
// @Summary Define an updated ball of an over
// Extra could be the following for an ball (wd, lb, b)
// Noball treated separately as either run or extra can happen
// Run could be the following for an ball (0-6)
type OverUpdateModel struct {
	Run    int                `json:"run" binding:"required"`
	Extra  string             `json:"extra"`
	NB     bool               `json:"nb" binding:"required"`
	Wicket WicketDetailsModel `json:"wicket"`
}

// WicketDetailsModel godoc
// @Summary Define wicket details
type WicketDetailsModel struct {
	BowlerID   string         `json:"bowlerID" binding:"required"`
	BatsmanID  string         `json:"batsmanID" binding:"required"`
	SupportID  string         `json:"supportID"`
	WicketType models.OutType `json:"wicketType" binding:"required"`
}

// CreateOverModel godoc
// @Summary Define create over model
type CreateOverModel struct {
	BowlerID string `json:"bowlerID" binding:"required"`
}

// NextBatsmanModel godoc
// @Summary Define add batsman model
type NextBatsmanModel struct {
	BatsmanID string `json:"batsmanID" binding:"required"`
}
