//Package domains defines db models
package domains

import (
	"cricket-scoreboard-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Over godoc
// @Summary Define Over model
type Over struct {
	ID         primitive.ObjectID
	OverNumber int
	Wide       int
	Noball     int
	Bye        int
	LB         int
	Zero       int
	One        int
	Two        int
	Three      int
	Four       int
	Five       int
	Six        int
	IsRunning  bool
	InningsID  primitive.ObjectID
	BowlerID   primitive.ObjectID
	Wickets    []Wicket
	Sequence   string
	Ball       int
}

// Wicket godoc
// @Summary Define Wicket model
type Wicket struct {
	BatsmanID primitive.ObjectID
	SupportID primitive.ObjectID
	OutType   models.OutType
}
