//Package domains defines db models
package domains

import (
	"cricket-scoreboard-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Match godoc
// @Summary Define Match model
type Match struct {
	ID           primitive.ObjectID
	Number       int
	MatchType    models.MatchType
	OverLimit    int
	Result       MatchResult
	SeriesID     primitive.ObjectID
	Participants []primitive.ObjectID
}

// MatchResult godoc
// @Summary Define MatchResult model
type MatchResult struct {
	Result        models.ResultType
	WinningTeamID primitive.ObjectID
	LosingTeamID  primitive.ObjectID
	WinLoseType   models.WinLoseType
}
