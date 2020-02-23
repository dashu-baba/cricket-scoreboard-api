//Package domains defines db models
package domains

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Team godoc
// @Summary Define Team domain model
type Team struct {
	ID      primitive.ObjectID
	Name    string
	Logo    string
	Players []Player
}
