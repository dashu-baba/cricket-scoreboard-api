//Package domains defines db models
package domains

import (
	"cricket-scoreboard-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Game godoc
// @Summary Define Game model
type Game struct {
	ID       primitive.ObjectID
	Name     string
	GameType models.GameType
	Teams    []Team
	Fixture  Fixture
}
