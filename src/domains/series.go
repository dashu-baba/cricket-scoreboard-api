//Package domains defines db models
package domains

import (
	"cricket-scoreboard-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Series godoc
// @Summary Define Series model
type Series struct {
	ID       primitive.ObjectID
	Name     string
	GameType models.GameType
	Teams    []SeriesParticipant
	Status   models.SeriesState
}

// SeriesParticipant godoc
// @Summary Define Series paticipant teams with squads.
type SeriesParticipant struct {
	TeamID       primitive.ObjectID
	SquadPlayers []primitive.ObjectID
}
