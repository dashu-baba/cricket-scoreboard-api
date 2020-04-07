package requestmodels

import (
	"cricket-scoreboard-api/src/models"
	"errors"

	"github.com/gin-gonic/gin"
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

// UpdateSeriesStatusModel godoc
// @Summary Define Series status update model
type UpdateSeriesStatusModel struct {
	Status models.SeriesState `json:"status" form:"status" xml:"status" binding:"required"`
}

//ValidateUpdateSeriesStatusRequests godoc
// @Summary validates the incoming requests
func ValidateUpdateSeriesStatusRequests(c *gin.Context) (UpdateSeriesStatusModel, error) {
	model := UpdateSeriesStatusModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}

//ValidateUpdateSquadRequests godoc
// @Summary validates the incoming requests
func ValidateUpdateSquadRequests(c *gin.Context) (UpdateSquadModel, error) {
	model := UpdateSquadModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}

//ValidateCreateMatchesRequests godoc
// @Summary validates the incoming requests
func ValidateCreateMatchesRequests(c *gin.Context) (MatchCreateModel, error) {
	model := MatchCreateModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	if len(model.Matches) < 1 {
		return model, errors.New("Minimum number of match")
	}

	for i := range model.Matches {
		if len(model.Matches[i].Participants) != 2 {
			return model, errors.New("Match should contains 2 teams")
		}
	}

	return model, nil
}

//ValidateCreateSeriesRequests method validates the requests payload of CreateSeries method
func ValidateCreateSeriesRequests(c *gin.Context) (SeriesCreateModel, error) {
	model := SeriesCreateModel{}
	model.Teams = []SeriesParticipantModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	if model.GameType == models.Bilateral && len(model.Teams) > 2 {
		return model, errors.New("Invalid data, bilateral series only have max 2 teams")
	}

	return model, nil
}

//ValidateAddTeamRequests method validates the requests payload of AddTeams and Teams method
func ValidateAddTeamRequests(c *gin.Context) (TeamsAddModel, error) {
	model := TeamsAddModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}

//ValidateRemoveTeamRequests method validates the requests payload of RemoveTeams and Teams method
func ValidateRemoveTeamRequests(c *gin.Context) (TeamsRemoveModel, error) {
	model := TeamsRemoveModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}
