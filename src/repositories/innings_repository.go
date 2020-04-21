//Package repositories defines the repository items.
package repositories

import (
	"context"
	"cricket-scoreboard-api/src/domains"
	"cricket-scoreboard-api/src/driver"
	"cricket-scoreboard-api/src/models"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const inningsCollectionName = "innings"

//InningsRepository defines the instance
type InningsRepository struct {
	DB *driver.DB
}

//NewInningsRepository returns a new InningsRepository.
func NewInningsRepository(DB *driver.DB) *InningsRepository {
	return &InningsRepository{
		DB: DB,
	}
}

//InsertMany godoc
//Insert a collection of innings and return the added collection
func (repo *InningsRepository) InsertMany(ctx context.Context, players []domains.Innings) {
	collections := repo.DB.Database.Collection(inningsCollectionName)
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

//GetLastInningsNumber godoc
//Get the last entry of the innings
func (repo *InningsRepository) GetLastInningsNumber(ctx context.Context) int {
	collections := repo.DB.Database.Collection(inningsCollectionName)

	option := options.FindOne()
	option.SetSort(bson.M{"number": -1})
	result := collections.FindOne(ctx, bson.M{}, option)
	if result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return 0
		}
		panic(result.Err())
	}

	innings := domains.Innings{}
	result.Decode(&innings)
	return innings.Number
}

//GetByID retrieves macth object from db by id
//and return that object.
func (repo *InningsRepository) GetByID(ctx context.Context, id string) domains.Innings {
	collections := repo.DB.Database.Collection(inningsCollectionName)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	findResult := collections.FindOne(ctx, bson.M{"id": objID})

	innings := domains.Innings{}
	err = findResult.Decode(&innings)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			panic(err)
		}
	}

	return innings
}

//GetCurrentInnings retrieves last active innings object from db by matchid
//and return that object.
func (repo *InningsRepository) GetCurrentInnings(ctx context.Context, matchid string) domains.Innings {
	collections := repo.DB.Database.Collection(inningsCollectionName)
	objID, err := primitive.ObjectIDFromHex(matchid)
	if err != nil {
		panic(err)
	}

	findResult := collections.FindOne(ctx, bson.M{"matchid": objID, "inningsstatus": models.OnGoing})

	innings := domains.Innings{}
	err = findResult.Decode(&innings)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			panic(err)
		}
	}

	return innings
}

//Update update a innings object into db
//and return that updated item.
func (repo *InningsRepository) Update(ctx context.Context, id string, updates map[string]interface{}) domains.Innings {
	collections := repo.DB.Database.Collection(inningsCollectionName)

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
