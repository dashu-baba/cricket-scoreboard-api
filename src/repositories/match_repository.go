//Package repositories defines the repository items.
package repositories

import (
	"context"
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/driver"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const matchCollectionName = "matches"

//MatchRepository defines the instance
type MatchRepository struct {
	DB *driver.DB
}

//NewMatchRepository returns a new MatchRepository.
func NewMatchRepository(DB *driver.DB) *MatchRepository {
	return &MatchRepository{
		DB: DB,
	}
}

//InsertMany godoc
//Insert a collection of match and retur the added collection
func (repo *MatchRepository) InsertMany(ctx context.Context, players []domains.Match) []domains.Match {
	collections := repo.DB.Database.Collection(matchCollectionName)
	items := []interface{}{}
	for _, value := range players {
		value.ID = primitive.NewObjectID()
		items = append(items, value)
	}
	_, err := collections.InsertMany(ctx, items)

	if err != nil {
		panic(err)
	}

	res := []domains.Match{}
	for _, val := range items {
		res = append(res, val.(domains.Match))
	}

	return res
}

//GetLastMatchNumber godoc
//Get the last entry of the match
func (repo *MatchRepository) GetLastMatchNumber(ctx context.Context) int {
	collections := repo.DB.Database.Collection(matchCollectionName)
	// groupStage := bson.D{{"$group", bson.D{{"max", bson.D{{"$max", "$number"}}}}}}
	// cursor, err := collections.Aggregate(ctx, mongo.Pipeline{groupStage})
	option := options.FindOne()
	option.SetSort(bson.M{"number": -1})
	result := collections.FindOne(ctx, bson.M{}, option)
	if result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return 0
		}
		panic(result.Err())
	}

	match := domains.Match{}
	result.Decode(&match)
	return match.Number
}

//GetAll retrieves all player objects from db
//by teamid
//and return that collection.
func (repo *MatchRepository) GetAll(teamID primitive.ObjectID) []domains.Player {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	collections := repo.DB.Database.Collection(matchCollectionName)
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
