//Package domains defines db models
package domains

import (
	"cricket-scoreboard-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Batting godoc
// @Summary Define Batting model
type Batting struct {
	ID          primitive.ObjectID
	Run         int
	Ball        int
	One         int
	Two         int
	Three       int
	Four        int
	Five        int
	Six         int
	IsInStrike  bool
	IsInCrease  bool
	OutType     models.OutType
	WicketBy    primitive.ObjectID
	SupportedBy primitive.ObjectID
	InningsID   primitive.ObjectID
	PlayerID    primitive.ObjectID
}
