//Package repositories defines the repository items.
package repositories

import (
	"context"
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/driver"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const battingCollectionName = "battings"

//BattingRepository defines the instance
type BattingRepository struct {
	DB *driver.DB
}

//NewBattingRepository returns a new BattingRepository.
func NewBattingRepository(DB *driver.DB) *BattingRepository {
	return &BattingRepository{
		DB: DB,
	}
}

//InsertMany godoc
//Insert a collection of batting and return the added collection
func (repo *BattingRepository) InsertMany(ctx context.Context, players []domains.Batting) {
	collections := repo.DB.Database.Collection(battingCollectionName)
	items := []interface{}{}
	for _, value := range players {
		value.ID = primitive.NewObjectID()
		items = append(items, value)
	}
	_, err := collections.InsertMany(ctx, items)

	if err != nil {
		panic(err)
	}
}

//GetByID retrieves macth object from db by id
//and return that object.
func (repo *BattingRepository) GetByID(ctx context.Context, id string) domains.Batting {
	collections := repo.DB.Database.Collection(battingCollectionName)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	findResult := collections.FindOne(ctx, bson.M{"id": objID})

	batting := domains.Batting{}
	err = findResult.Decode(&batting)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			panic(err)
		}
	}

	return batting
}

//GetCurrentBatsman retrieves current batsman in crease
//and return the collection.
func (repo *BattingRepository) GetCurrentBatsman(ctx context.Context, inningsID string) []domains.Batting {
	collections := repo.DB.Database.Collection(battingCollectionName)
	objID, err := primitive.ObjectIDFromHex(inningsID)
	if err != nil {
		panic(err)
	}

	cursor, err := collections.Find(ctx, bson.M{"inningsid": objID, "isincrease": true})

	battings := []domains.Batting{}
	if err != nil {
		panic(err)
	}

	for cursor.Next(ctx) {
		batting := domains.Batting{}
		cursor.Decode(&batting)
		battings = append(battings, batting)
	}

	return battings
}

//Update update a batting object into db
//and return that updated item.
func (repo *BattingRepository) Update(ctx context.Context, id string, updates map[string]interface{}) domains.Batting {
	collections := repo.DB.Database.Collection(battingCollectionName)

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

	return repo.GetByID(ctx, id)
}
