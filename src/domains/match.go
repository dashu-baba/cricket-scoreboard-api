//Package domains defines db models
package domains

import (
	"cricket-scoreboard-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Match godoc
// @Summary Define Match model
type Match struct {
	ID        primitive.ObjectID
	Number    int
	MatchType models.MatchType
	Over      int
}
