//Package domains defines db models
package domains

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fixture godoc
// @Summary Define Fixture model
type Fixture struct {
	ID      primitive.ObjectID
	Matches []Match
	GameID  primitive.ObjectID
}
