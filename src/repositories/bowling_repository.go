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

const bowlingCollectionName = "bowlings"

//BowlingRepository defines the instance
type BowlingRepository struct {
	DB *driver.DB
}

//NewBowlingRepository returns a new BowlingRepository.
func NewBowlingRepository(DB *driver.DB) *BowlingRepository {
	return &BowlingRepository{
		DB: DB,
	}
}

//InsertMany godoc
//Insert a collection of bowling and return the added collection
func (repo *BowlingRepository) InsertMany(ctx context.Context, bowlers []domains.Bowling) {
	collections := repo.DB.Database.Collection(bowlingCollectionName)
	items := []interface{}{}
	for _, value := range bowlers {
		value.ID = primitive.NewObjectID()
		items = append(items, value)
	}
	_, err := collections.InsertMany(ctx, items)

	if err != nil {
		panic(err)
	}
}

//GetCurrentBowler retrieves current bowler in crease
//and return the collection.
func (repo *BowlingRepository) GetCurrentBowler(ctx context.Context, inningsID string) domains.Bowling {
	collections := repo.DB.Database.Collection(bowlingCollectionName)
	objID, err := primitive.ObjectIDFromHex(inningsID)
	if err != nil {
		panic(err)
	}

	findResult := collections.FindOne(ctx, bson.M{"inningsid": objID, "iscurrent": true})

	bowling := domains.Bowling{}
	err = findResult.Decode(&bowling)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			panic(err)
		}
	}

	return bowling
}

//GetByID retrieves macth object from db by id
//and return that object.
func (repo *BowlingRepository) GetByID(ctx context.Context, id string) domains.Bowling {
	collections := repo.DB.Database.Collection(bowlingCollectionName)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	findResult := collections.FindOne(ctx, bson.M{"id": objID})

	bowling := domains.Bowling{}
	err = findResult.Decode(&bowling)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			panic(err)
		}
	}

	return bowling
}

//Update update a bowling object into db
//and return that updated item.
func (repo *BowlingRepository) Update(ctx context.Context, id string, updates map[string]interface{}) domains.Bowling {
	collections := repo.DB.Database.Collection(bowlingCollectionName)

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

//GetAllByIds godoc
//Find collection by a collection of id and returns the collection
func (repo *BowlingRepository) GetAllByIds(ctx context.Context, ids []string) []domains.Bowling {
	oids := []primitive.ObjectID{}
	for _, val := range ids {
		oid, err := primitive.ObjectIDFromHex(val)
		if err != nil {
			panic(err)
		}
		oids = append(oids, oid)
	}

	collections := repo.DB.Database.Collection(bowlingCollectionName)

	cursor, err := collections.Find(ctx, bson.M{"id": bson.M{"$in": oids}})

	if err != nil {
		panic(err)
	}

	bowlers := []domains.Bowling{}
	for cursor.Next(ctx) {
		player := domains.Bowling{}
		cursor.Decode(&player)
		bowlers = append(bowlers, player)
	}

	return bowlers
}
