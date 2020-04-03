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

const seriesCollectionName = "series"

//SeriesRepository defines the instance
type SeriesRepository struct {
	DB *driver.DB
}

//NewSeriesRepository returns a new SeriesRepository.
func NewSeriesRepository(DB *driver.DB) *SeriesRepository {
	return &SeriesRepository{
		DB: DB,
	}
}

//Insert insert a series object into db
//and return that inserted item.
func (repo *SeriesRepository) Insert(ctx context.Context, series domains.Series) domains.Series {
	collections := repo.DB.Database.Collection(seriesCollectionName)
	series.ID = primitive.NewObjectID()
	_, err := collections.InsertOne(ctx, series)

	if err != nil {
		panic(err)
	}

	return series
}

//Update update a player object into db
//and return that updated item.
func (repo *SeriesRepository) Update(ctx context.Context, id string, updates map[string]interface{}) domains.Series {
	collections := repo.DB.Database.Collection(seriesCollectionName)

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

//toDoc converts object to bson document
func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

//GetAll retrieves all series objects from db
//and return that collection.
func (repo *SeriesRepository) GetAll(ctx context.Context) []domains.Series {
	collections := repo.DB.Database.Collection(seriesCollectionName)
	cursor, err := collections.Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}

	seriess := []domains.Series{}
	for cursor.Next(ctx) {
		series := domains.Series{}
		cursor.Decode(&series)
		seriess = append(seriess, series)
	}

	return seriess
}

//GetAllByIds retrieves series objects from db by ids
//and return that collection.
func (repo *SeriesRepository) GetAllByIds(ctx context.Context, ids []string) []domains.Series {
	collections := repo.DB.Database.Collection(seriesCollectionName)

	oids := []primitive.ObjectID{}
	for _, val := range ids {
		oid, err := primitive.ObjectIDFromHex(val)
		if err != nil {
			panic(err)
		}
		oids = append(oids, oid)
	}
	cursor, err := collections.Find(ctx, bson.M{"id": bson.M{"in": oids}})

	if err != nil {
		panic(err)
	}

	seriess := []domains.Series{}
	for cursor.Next(ctx) {
		series := domains.Series{}
		cursor.Decode(&series)
		seriess = append(seriess, series)
	}

	return seriess
}

//GetByID retrieves series object from db by id
//and return that object.
func (repo *SeriesRepository) GetByID(ctx context.Context, id string) domains.Series {
	collections := repo.DB.Database.Collection(seriesCollectionName)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	findResult := collections.FindOne(ctx, bson.M{"id": objID})

	if err := findResult.Err(); err != nil {
		panic(err)
	}

	series := domains.Series{}
	err = findResult.Decode(&series)
	if err != nil {
		panic(err)
	}

	return series
}
