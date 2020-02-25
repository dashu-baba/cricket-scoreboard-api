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

//InsertMany insert a player object into db
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

	res := []domains.Player{}
	for _, val := range items {
		res = append(res, val.(domains.Player))
	}

	return res
}

//GetAll retrieves all player objects from db
//by teamid
//and return that collection.
func (repo *PlayerRepository) GetAll(teamID primitive.ObjectID) []domains.Player {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
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
