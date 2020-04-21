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

const gameCollectionName = "games"

//GameRepository defines the instance
type GameRepository struct {
	DB *driver.DB
}

//NewGameRepository returns a new GameRepository.
func NewGameRepository(DB *driver.DB) *GameRepository {
	return &GameRepository{
		DB: DB,
	}
}

//Insert insert a game object into db
//and return that inserted item.
func (repo *GameRepository) Insert(game domains.Game) domains.Game {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collections := repo.DB.Database.Collection(gameCollectionName)

	game.ID = primitive.NewObjectID()
	_, err := collections.InsertOne(ctx, game)

	if err != nil {
		panic(err)
	}

	return game
}

//Update update a game object
//and return it.
func (repo *GameRepository) Update(game domains.Game, toBeUpdated map[string]interface{}) domains.Game {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collections := repo.DB.Database.Collection(gameCollectionName)
	filter := bson.M{"id": game.ID}
	update := bson.M{"$set": toBeUpdated}

	_, err := collections.UpdateOne(
		ctx,
		filter,
		update,
	)

	if err != nil {
		panic(err)
	}

	return repo.GetByID(game.ID)
}

//GetByID get single item by id and returns it.
func (repo *GameRepository) GetByID(id primitive.ObjectID) domains.Game {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collections := repo.DB.Database.Collection(gameCollectionName)

	res := collections.FindOne(ctx, bson.M{"id": id})

	if res.Err() != nil {
		panic(res.Err())
	}

	game := domains.Game{}
	res.Decode(&game)

	return game
}
