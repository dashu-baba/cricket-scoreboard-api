package validators

import (
	"cricket-scoreboard-api/src/models"
	"cricket-scoreboard-api/src/requestmodels"
	"errors"

	"github.com/gin-gonic/gin"
)

//ValidateUpdateMatchPlayingSquadRequests godoc
// @Summary validates the incoming requests
func ValidateUpdateMatchPlayingSquadRequests(c *gin.Context) (requestmodels.MatchPlayingSquadModel, error) {
	model := requestmodels.MatchPlayingSquadModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	if len(model.Players) < 1 {
		return model, errors.New("Please enter the minimum number of playing squads")
	}

	for _, val := range model.Players {
		for _, val1 := range model.Extras {
			if val == val1 {
				return model, errors.New("Same player cannot be in playing squad and reserve")
			}
		}
	}

	return model, nil
}

//ValidateUpdateSeriesStatusRequests godoc
// @Summary validates the incoming requests
func ValidateUpdateSeriesStatusRequests(c *gin.Context) (requestmodels.UpdateSeriesStatusModel, error) {
	model := requestmodels.UpdateSeriesStatusModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}

//ValidateUpdateSquadRequests godoc
// @Summary validates the incoming requests
func ValidateUpdateSquadRequests(c *gin.Context) (requestmodels.UpdateSquadModel, error) {
	model := requestmodels.UpdateSquadModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}

//ValidateCreateMatchesRequests godoc
// @Summary validates the incoming requests
func ValidateCreateMatchesRequests(c *gin.Context) (requestmodels.MatchCreateModel, error) {
	model := requestmodels.MatchCreateModel{}
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
func ValidateCreateSeriesRequests(c *gin.Context) (requestmodels.SeriesCreateModel, error) {
	model := requestmodels.SeriesCreateModel{}
	model.Teams = []requestmodels.SeriesParticipantModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	if model.GameType == models.Bilateral && len(model.Teams) > 2 {
		return model, errors.New("Invalid data, bilateral series only have max 2 teams")
	}

	return model, nil
}

//ValidateAddTeamRequests method validates the requests payload of AddTeams and Teams method
func ValidateAddTeamRequests(c *gin.Context) (requestmodels.TeamsAddModel, error) {
	model := requestmodels.TeamsAddModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}

//ValidateRemoveTeamRequests method validates the requests payload of RemoveTeams and Teams method
func ValidateRemoveTeamRequests(c *gin.Context) (requestmodels.TeamsRemoveModel, error) {
	model := requestmodels.TeamsRemoveModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}
