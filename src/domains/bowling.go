//Package domains defines db models
package domains

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Bowling godoc
// @Summary Define Bowling model
type Bowling struct {
	ID        primitive.ObjectID
	Over      int
	Wicket    int
	Wide      int
	Noball    int
	Three     int
	Four      int
	Five      int
	Six       int
	IsCurrent bool
	One       int
	InningsID primitive.ObjectID
	PlayerID  primitive.ObjectID
	Overs     []Over
}
