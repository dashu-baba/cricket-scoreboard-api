//Package domains defines db models
package domains

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Game godoc
// @Summary Define Game model
type Game struct {
	ID      primitive.ObjectID
	MatchID primitive.ObjectID
	Teams   []Team
}
