//Package repositories defines the repository items.
package repositories

import (
	"context"
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/driver"
	"fmt"

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

//Update update a team object into db
//and return that updated item.
func (repo *PlayerRepository) Update(ctx context.Context, id string, updates map[string]interface{}) {
	collections := repo.DB.Database.Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"id": objID}
	updatedValue := bson.M{}
	updatedValue = updates

	update := bson.M{"$set": updatedValue}
	updateResult, err := collections.UpdateOne(
		ctx,
		filter,
		update,
	)

	fmt.Println(updateResult)

	if err != nil {
		panic(err)
	}
}

//InsertMany insert a list of player objects into db
//and return that inserted items.
func (repo *PlayerRepository) InsertMany(ctx context.Context, players []domains.Player) []domains.Player {
	if len(players) <= 0 {
		return []domains.Player{}
	}

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

	return repo.GetAll(ctx, players[0].TeamID.Hex())
}

//Insert insert a player object into db
//and return that inserted item.
func (repo *PlayerRepository) Insert(ctx context.Context, player domains.Player) domains.Player {
	collections := repo.DB.Database.Collection(collectionName)

	player.ID = primitive.NewObjectID()
	_, err := collections.InsertOne(ctx, player)

	if err != nil {
		panic(err)
	}

	return player
}

//Remove removes a player object from db
func (repo *PlayerRepository) Remove(ctx context.Context, id primitive.ObjectID) {
	collections := repo.DB.Database.Collection(collectionName)

	_, err := collections.DeleteOne(ctx, bson.M{"id": id})

	if err != nil {
		panic(err)
	}
}

//GetAll retrieves all player objects from db
//by teamid
//and return that collection.
func (repo *PlayerRepository) GetAll(ctx context.Context, teamID string) []domains.Player {
	objID, err := primitive.ObjectIDFromHex(teamID)
	if err != nil {
		panic(err)
	}

	collections := repo.DB.Database.Collection(collectionName)
	cursor, err := collections.Find(ctx, bson.M{"teamid": objID})

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

//GetAllByIds godoc
//Find player collection by a collection of id and returns the collection
func (repo *PlayerRepository) GetAllByIds(ctx context.Context, ids []string) []domains.Player {
	oids := []primitive.ObjectID{}
	for _, val := range ids {
		oid, err := primitive.ObjectIDFromHex(val)
		if err != nil {
			panic(err)
		}
		oids = append(oids, oid)
	}

	collections := repo.DB.Database.Collection(collectionName)

	cursor, err := collections.Find(ctx, bson.M{"id": bson.M{"$in": oids}})

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
