package requestmodels

import (
	"cricket-scoreboard-api/src/models"
)

// SeriesCreateModel godoc
// @Summary Define Game create model
type SeriesCreateModel struct {
	Name     string                   `json:"name" form:"name" xml:"name" binding:"required"`
	GameType models.GameType          `json:"gameType" form:"gametype" xml:"gameType" binding:"required"`
	Teams    []SeriesParticipantModel `json:"teams" form:"teams" xml:"teams" binding:"required"`
}

// TeamsAddModel godoc
// @Summary Define team add model
type TeamsAddModel struct {
	Teams []SeriesParticipantModel `json:"teams" form:"teams" xml:"teams" binding:"required"`
}

// TeamsRemoveModel godoc
// @Summary Define team remove model
type TeamsRemoveModel struct {
	Teams []string `json:"teams" form:"teams" xml:"teams" binding:"required"`
}

// SeriesParticipantModel godoc
// @Summary Define participant details of series
type SeriesParticipantModel struct {
	TeamID       string   `json:"teamId" form:"teamid" xml:"teamId" binding:"required"`
	SquadPlayers []string `json:"squadPlayers" form:"squadplayers" xml:"squadPlayers" binding:"required"`
}

// MatchCreateModel godoc
// @Summary Define Match create model
type MatchCreateModel struct {
	Matches []Match `json:"matches" form:"matches" xml:"matches" binding:"required"`
}

// Match godoc
// @Summary Define Match
type Match struct {
	MatchType    models.MatchType `json:"matchType" form:"matchtype" xml:"matchType" binding:"required"`
	OverLimit    int              `json:"overLimit" form:"overlimit" xml:"overLimit" binding:"required"`
	Participants []string         `json:"participants" form:"participants" xml:"participants" binding:"required"`
}

// UpdateSquadModel godoc
// @Summary Define Team squad update model
type UpdateSquadModel struct {
	AddedPlayer   []string `json:"addedPlayer" form:"addedplayer" xml:"addedPlayer"`
	RemovedPlayer []string `json:"removedPlayer" form:"removedplayer" xml:"removedPlayer"`
}

// MatchPlayingSquadModel godoc
// @Summary Define match squad update model
type MatchPlayingSquadModel struct {
	TeamID  string   `json:"teamId" form:"teamid" xml:"teamId" binding:"required"`
	Players []string `json:"players" form:"players" xml:"players"`
	Extras  []string `json:"extras" form:"extras" xml:"extras"`
}

// UpdateSeriesStatusModel godoc
// @Summary Define Series status update model
type UpdateSeriesStatusModel struct {
	Status models.SeriesState `json:"status" form:"status" xml:"status" binding:"required"`
}
