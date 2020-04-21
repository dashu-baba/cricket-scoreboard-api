//Package domains defines db models
package domains

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Bowling godoc
// @Summary Define Bowling model
type Bowling struct {
	ID        primitive.ObjectID
	IsCurrent bool
	InningsID primitive.ObjectID
	PlayerID  primitive.ObjectID
	Overs     []Over
	Wickets   int
}
