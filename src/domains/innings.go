//Package domains defines db models
package domains

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Innings godoc
// @Summary Define Innings model
type Innings struct {
	ID          primitive.ObjectID
	Number      int
	OverLimit   int
	OverPlayed  int
	MatchID     primitive.ObjectID
	BattingTeam Team
	BowlingTeam Team
}
