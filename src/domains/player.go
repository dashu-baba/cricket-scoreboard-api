//Package domains defines db models
package domains

import (
	"cricket-scoreboard-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Player godoc
// @Summary Define Player model
type Player struct {
	ID         primitive.ObjectID
	Name       string
	PlayerType models.PlayerType
	TeamID     primitive.ObjectID
}
