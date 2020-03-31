//Package repositories defines the repository items.
package repositories

import (
	"context"
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/driver"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName = "players"

//PlayerRepository defines the instance
type PlayerRepository struct {
	DB *driver.DB
}

//NewPlayerRepository returns a new PlayerRepository.
func NewPlayerRepository(DB *driver.DB) *PlayerRepository {
	return &PlayerRepository{
		DB: DB,
	}
}

//InsertMany insert a list of player objects into db
//and return that inserted items.
func (repo *PlayerRepository) InsertMany(players []domains.Player) []domains.Player {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collections := repo.DB.Database.Collection(collectionName)
	items := []interface{}{}
	for _, value := range players {
		value.ID = primitive.NewObjectID()
		items = append(items, value)
	}
	_, err := collections.InsertMany(ctx, items)

	if err != nil {
		panic(err)
	}

	return repo.GetAll(players[0].TeamID)
}

//Insert insert a player object into db
//and return that inserted item.
func (repo *PlayerRepository) Insert(player domains.Player) domains.Player {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collections := repo.DB.Database.Collection(collectionName)

	player.ID = primitive.NewObjectID()
	_, err := collections.InsertOne(ctx, player)

	if err != nil {
		panic(err)
	}

	return player
}

//Remove removes a player object from db
func (repo *PlayerRepository) Remove(id primitive.ObjectID) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collections := repo.DB.Database.Collection(collectionName)

	_, err := collections.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		panic(err)
	}
}

//GetAll retrieves all player objects from db
//by teamid
//and return that collection.
func (repo *PlayerRepository) GetAll(teamID primitive.ObjectID) []domains.Player {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collections := repo.DB.Database.Collection(collectionName)
	cursor, err := collections.Find(ctx, bson.M{"teamID": teamID})

	if err != nil {
		panic(err)
	}

	players := []domains.Player{}
	for cursor.Next(ctx) {
		player := domains.Player{}
		cursor.Decode(&player)
		players = append(players, player)
	}

	return players
}
