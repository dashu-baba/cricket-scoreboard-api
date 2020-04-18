//Package domains defines db models
package domains

import (
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
	One        int
	Two        int
	Three      int
	Four       int
	Five       int
	Six        int
	IsRunning  bool
	InningsID  primitive.ObjectID
	BowlerID   primitive.ObjectID
	Wickets    []Batting
	Squence    string
}
