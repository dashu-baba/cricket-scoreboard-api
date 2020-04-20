//Package domains defines db models
package domains

import (
	"cricket-scoreboard-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Innings godoc
// @Summary Define Innings model
type Innings struct {
	ID            primitive.ObjectID
	Number        int
	OverLimit     int
	OverPlayed    float64
	MatchID       primitive.ObjectID
	BattingTeamID primitive.ObjectID
	BowlingTeamID primitive.ObjectID
	TossResult    primitive.ObjectID
	InningsStatus models.SeriesState
	Run           int
	Wicket        int
	WicketLimit   int
	Target        int
}
