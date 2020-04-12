//Package driver represents the startup setup works
package driver

import (
	"context"
	"cricket-scoreboard-api/src/models"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//DB represent the database connection
type DB struct {
	Database *mongo.Database
	Context  context.Context
}

//ConnectDb creates a connection to database
func ConnectDb() *DB {
	configuration := models.Configuration
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configuration.Db.EndPoint))

	if err != nil {
		panic(err)
	}

	return &DB{
		Database: client.Database(configuration.Db.Name),
		Context:  ctx,
	}
}
