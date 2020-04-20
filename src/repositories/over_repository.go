//Package repositories defines the repository items.
package repositories

import (
	"context"
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/driver"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const overCollectionName = "overs"

//OverRepository defines the instance
type OverRepository struct {
	DB *driver.DB
}

//NewOverRepository returns a new OverRepository.
func NewOverRepository(DB *driver.DB) *OverRepository {
	return &OverRepository{
		DB: DB,
	}
}

//InsertMany godoc
//Insert a collection of over and return the added collection
func (repo *OverRepository) InsertMany(ctx context.Context, players []domains.Over) {
	collections := repo.DB.Database.Collection(overCollectionName)
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

//GetLastOverNumber godoc
//Get the last entry of the over
func (repo *OverRepository) GetLastOverNumber(ctx context.Context, inningsid string) int {
	collections := repo.DB.Database.Collection(overCollectionName)
	objID, err := primitive.ObjectIDFromHex(inningsid)
	if err != nil {
		panic(err)
	}
	option := options.FindOne()
	option.SetSort(bson.M{"overnumber": -1})
	result := collections.FindOne(ctx, bson.M{"inningsid": objID}, option)
	if result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return 0
		}
		panic(result.Err())
	}

	over := domains.Over{}
	result.Decode(&over)
	return over.OverNumber
}

//GetByID retrieves macth object from db by id
//and return that object.
func (repo *OverRepository) GetByID(ctx context.Context, id string) domains.Over {
	collections := repo.DB.Database.Collection(overCollectionName)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	findResult := collections.FindOne(ctx, bson.M{"id": objID})

	over := domains.Over{}
	err = findResult.Decode(&over)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domains.Over{}
		} else {
			panic(err)
		}
	}

	return over
}

//HasAnyRunningOver checks an running over exists by innings id
//and return bool.
func (repo *OverRepository) HasAnyRunningOver(ctx context.Context, inningsid string) bool {
	collections := repo.DB.Database.Collection(overCollectionName)
	objID, err := primitive.ObjectIDFromHex(inningsid)
	if err != nil {
		panic(err)
	}

	result := collections.FindOne(ctx, bson.M{"inningsid": objID, "isrunning": true})

	if result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return false
		}
		panic(result.Err())
	}

	return true
}

//Update update a over object into db
//and return that updated item.
func (repo *OverRepository) Update(ctx context.Context, id string, updates map[string]interface{}) domains.Over {
	collections := repo.DB.Database.Collection(overCollectionName)

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
